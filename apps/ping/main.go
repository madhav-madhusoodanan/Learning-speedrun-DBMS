package ping

import (

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	router.Use(cors.Default())
	router.GET("/", Pong)
	router.GET("/age/:name", AgePredict)
}
