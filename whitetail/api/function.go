package api

import (
	"net/http"
	"whitetail/function"
	"whitetail/operation"
	"whitetail/utils"

	"github.com/gin-gonic/gin"
)

func RunFunction(ctx *gin.Context) {
	name := ctx.Param("name")

	var data map[string]interface{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}

	if err := function.Run(operation.Operations.Functions[name], data); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
