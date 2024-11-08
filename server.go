package main

import (
	"github.com/gin-gonic/gin"
	"rampx/backend/apps/ping"
	"rampx/backend/apps/swaps"
)

func main() {
	r := gin.Default()

	swaps.Setup(r.Group("/swap"))
	ping.Setup(r.Group("/ping"))

	r.Run()

}
