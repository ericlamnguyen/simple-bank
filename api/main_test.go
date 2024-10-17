package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// Setup code (run before all tests)
	gin.SetMode(gin.TestMode)

	// Run the tests
	exitCode := m.Run()

	// Teardown code (run after all tests)

	// Exit with the code returned from m.Run()
	os.Exit(exitCode)
}
