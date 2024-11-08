package main

import (
	"rampx/backend/apps/ping"
	"rampx/backend/apps/swaps"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	swaps.Setup(r.Group("/swap"))
	ping.Setup(r.Group("/ping"))

	r.Run()

}
