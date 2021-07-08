package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAuth registers the auth endpoint
func RegisterAuth(router *gin.RouterGroup) {
	fmt.Println("in auth")
	auth := router.Group("/auth")

	auth.GET("/name", handleName)
}

func handleName(c *gin.Context) {
	fmt.Println(c.Params) 
	fmt.Println(c.GetHeader("Authorization"))
	c.JSON(http.StatusAccepted, "[]")
}
