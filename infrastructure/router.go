package infrastructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//GinRouter -> Gin Router
type GinRouter struct {
	Gin *gin.Engine
}

//NewGinRouter all the routes are defined here
func NewGinRouter() GinRouter {

	httpRouter := gin.Default()

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})
	return GinRouter{
		Gin: httpRouter,
	}

}
