package api

import (
	db "github.com/ericlamnguyen/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for banking services
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routings
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	// Add routes to router
	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	// Attach router to server
	server.router = router

	return server
}

// Run the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Function to convert error into map[string]any
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
