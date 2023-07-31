package api

import (
	db "github.com/atanashristov/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves all API requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// Creates new server instance
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Run the http server on the provided address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
