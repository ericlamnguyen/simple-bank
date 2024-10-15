package api

import (
	"bytes"
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

// Test getAccount using mocking
func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	// create mock object for Store interface
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedStore := mockdb.NewMockStore(ctrl)
	mockedStore.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// start test server and send request
	testServer := NewServer(mockedStore)

	// create new request
	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	// Create a ResponseRecorder to record the response
	// httptest.ResponseRecorder implements the http.ResponseWriter
	rr := httptest.NewRecorder()

	// send request through router and record the response in recorder
	testServer.router.ServeHTTP(rr, request)

	// check the response
	require.Equal(t, http.StatusOK, rr.Code)
	requireBodyMatchAccount(t, rr.Body, account)
}

// Create a random account for testing
func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

// verify response body matches with account object
func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}
