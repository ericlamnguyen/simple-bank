package api

import (
	"os"
	"testing"
	"time"

	db "github.com/ericlamnguyen/simple-bank/db/sqlc"
	"github.com/ericlamnguyen/simple-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	testServer, err := NewServer(config, store)
	require.NoError(t, err)

	return testServer
}

func TestMain(m *testing.M) {
	// Setup code (run before all tests)
	gin.SetMode(gin.TestMode)

	// Run the tests
	exitCode := m.Run()

	// Teardown code (run after all tests)

	// Exit with the code returned from m.Run()
	os.Exit(exitCode)
}
