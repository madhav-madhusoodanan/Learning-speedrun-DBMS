package swaps

import (
	"rampx/backend/apps/swaps/functions"
	"rampx/backend/apps/swaps/utils"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	utils.InitializeContracts()
	utils.InitializeRPC()

	router.POST("/execute", functions.HandleCrossChainSwap)
}
