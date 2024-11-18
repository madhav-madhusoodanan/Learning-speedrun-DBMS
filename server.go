package main

import (
	"rampx/backend/apps/ping"
	"rampx/backend/apps/swaps"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())
	swaps.Setup(r.Group("/swap"))
	ping.Setup(r.Group("/ping"))

	r.Run()

}
