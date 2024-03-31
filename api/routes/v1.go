package routes

import (
	handler "etf-data-extractor/api/handler"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.GET("/checkHealth", handler.CheckHealth)
	router.GET("/get-crypto-data", handler.FetchCryptoData())
}
