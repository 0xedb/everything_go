package app

import (
	"log"

	"github.com/0xedb/everything_go/authentication/auth"
	"github.com/gin-gonic/gin"
)

const defaultPort string = ":2021"

var server *gin.Engine

// GetInstance returns the server instance
func GetInstance() *gin.RouterGroup {
	if server == nil {
		server = gin.Default()
	}

	return &server.RouterGroup
}

// StartServer starts the app server
func StartServer() {
	if server == nil {
		GetInstance()
	}

	auth.RegisterAuth(GetInstance())

	log.Println("starting server on http://localhost" + defaultPort)
	server.Run(defaultPort)
}
