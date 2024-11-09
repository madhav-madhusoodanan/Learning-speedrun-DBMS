package swaps

import (
	"rampx/backend/apps/swaps/functions"
	"rampx/backend/apps/swaps/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	utils.InitializeContracts()
	utils.InitializeRPC()

	router.Use(cors.Default())
	router.POST("/execute", functions.HandleCrossChainSwap)
}
