package swaps

import (
	"rampx/backend/internal/apps/swaps/internal/functions"
	"rampx/backend/internal/apps/swaps/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	utils.InitializeContracts()
	utils.InitializeRPC()

	router.Use(cors.Default())
	router.POST("/execute", functions.HandleCrossChainSwap)
}
