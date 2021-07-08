package app

import (
	"log"

	"github.com/0xedb/everything_go/authentication/auth"
	"github.com/gin-gonic/gin"
)

const defaultPort string = ":2021"

var server *gin.Engine

func init() {
	server = gin.Default()

	// register routes
	auth.RegisterAuth(GetInstance())
}

// GetInstance returns the server instance
func GetInstance() *gin.RouterGroup {
	return &server.RouterGroup
}

// StartServer starts the app server
func StartServer() {
	if server == nil {
		GetInstance()
	}

	log.Println("starting server on http://localhost" + defaultPort)
	server.Run(defaultPort)
}
