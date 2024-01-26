package api

import (
	"net/http"
	"whitetail/basestation"
	"whitetail/utils"

	"github.com/gin-gonic/gin"
)

func UpdateBaseStation(ctx *gin.Context) {
	observer := ctx.Param("observer")
	name := ctx.Param("stream")

	var data map[string]string
	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}

	if err := basestation.ReceiveData(name, observer, data["data"]); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetBaseStation(ctx *gin.Context) {
	observer := ctx.Param("observer")
	name := ctx.Param("stream")

	var filter interface{}
	if err := ctx.BindJSON(filter); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}

	data, err := basestation.GetData(name, observer, filter)
	if err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, data)
}
