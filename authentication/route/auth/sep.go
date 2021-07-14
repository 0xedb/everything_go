package auth

import (
	"net/http"

	"github.com/0xedb/everything_go/authentication/route"
	"github.com/gin-gonic/gin"
)

// Register register auth enpoint
func (r auth) Register(router *gin.RouterGroup) {
	auth := router.Group(root)

	auth.GET(ext, handleExt)
	auth.GET(name, handleName)
	auth.GET(options, handleOption)
}

func handleName(c route.Ctx) {
	c.JSON(http.StatusOK, gin.H{
		"name": "bruno",
		"age":  0xedb,
		"fee":  0xdead,
	})
}

func handleExt(c route.Ctx) {
	c.String(http.StatusOK, "hello")
}
