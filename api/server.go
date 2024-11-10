package api

import (
	"fmt"

	db "github.com/ericlamnguyen/simple-bank/db/sqlc"
	"github.com/ericlamnguyen/simple-bank/token"
	"github.com/ericlamnguyen/simple-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for banking services
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routings
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// Create a new PASETO token maker to create and verify PASETO tokens
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	// Create Server object
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// Register additional parameter validator with gin
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency) // ensure only supported currency is allowed
	}

	// Attach router to server
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	// Add route handlers
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("users/login", server.loginUser)

	// The following endpoints require authMiddleware
	// router := router.Group("/").Use(authMiddleware(server.tokenMaker))

	router.POST("/accounts", authMiddleware(server.tokenMaker), server.createAccount)
	router.GET("/accounts/:id", authMiddleware(server.tokenMaker), server.getAccount)
	router.GET("/accounts", authMiddleware(server.tokenMaker), server.listAccount)

	router.POST("/transfers", authMiddleware(server.tokenMaker), server.createTransfer)

	// Attach router to server
	server.router = router
}

// Run the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Function to convert error into map[string]any
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
