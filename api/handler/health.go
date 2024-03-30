package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealth(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "Health-OK")
}
