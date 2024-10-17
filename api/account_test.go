package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/ericlamnguyen/simple-bank/db/mock"
	db "github.com/ericlamnguyen/simple-bank/db/sqlc"
	"github.com/ericlamnguyen/simple-bank/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// Test Server.getAccount method using mocking
func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	// create table-driven test suite
	testCases := []struct {
		testName      string
		accountID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, rr *httptest.ResponseRecorder)
	}{
		// test 1 - successful request
		{
			testName:  "OK",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rr.Code)
				requireBodyMatchAccount(t, rr.Body, account)
			},
		},
		// test 2 - requested account not found in db
		{
			testName:  "NotFound",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, rr.Code)
				requireBodyMatchAccount(t, rr.Body, db.Account{})
			},
		},
		// test 3 - server internal error encountered
		{
			testName:  "InternalError",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rr.Code)
				requireBodyMatchAccount(t, rr.Body, db.Account{})
			},
		},
		// test 4 - invalid parameters
		{
			testName:  "InvalidID",
			accountID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rr.Code)
			},
		},
	}

	// run the test suite
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.testName, func(t *testing.T) {
			// Create a new mock controller to manage the lifecycle of mocks during the test
			ctrl := gomock.NewController(t)
			// Ensure that the mocks are validated and any expected calls are checked after the test
			defer ctrl.Finish()

			// create mock object for Store interface
			mockedStore := mockdb.NewMockStore(ctrl)
			tc.buildStubs(mockedStore)

			// start test server
			testServer := NewServer(mockedStore)

			// create new request
			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Create a ResponseRecorder to record the response
			// httptest.ResponseRecorder implements the http.ResponseWriter
			rr := httptest.NewRecorder()

			// send request through router and record the response in recorder
			testServer.router.ServeHTTP(rr, request)

			// check the response
			tc.checkResponse(t, rr)
		})
	}
}

// helper function to Create a random account for testing
func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

// helper function to verify response body matches with account object
func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}
