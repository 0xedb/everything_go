package route

import "github.com/gin-gonic/gin"

// Path is a path
type Path interface {
	Register(*gin.RouterGroup)
}

// Ctx is a gin context
type Ctx = *gin.Context
