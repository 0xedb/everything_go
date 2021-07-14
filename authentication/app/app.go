package app

import (
	"log"

	"github.com/0xedb/everything_go/authentication/route"
	"github.com/0xedb/everything_go/authentication/route/auth"
	"github.com/gin-gonic/gin"
)

const defaultPort string = ":2021"

var server *gin.Engine

func init() {
	server = gin.Default()

	endpoints := []route.Path{auth.Instance}

	registerRoutes(endpoints)
}

func registerRoutes(routes []route.Path) {
	instance := GetInstance()

	for _, route := range routes {
		route.Register(instance)
	}
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
