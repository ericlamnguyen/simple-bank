package api

import (
	db "github.com/ericlamnguyen/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for banking services
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routings
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	// Add route handlers
	router := gin.Default()

	// Register additional parameter validator with gin
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency) // ensure only supported currency is allowed
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	router.POST("/users", server.createUser)

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
