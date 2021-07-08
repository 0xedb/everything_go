package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const authEnpoint string = "/auth"

// RegisterAuth registers the auth endpoint
func RegisterAuth(router *gin.RouterGroup) {
	fmt.Println("in auth")
	auth := router.Group(authEnpoint)

	auth.GET("/name", handleName)
}

func handleName(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"name":"bruno",
		"age": 0xedb,
		"fee": 0xdead,
	})
}
