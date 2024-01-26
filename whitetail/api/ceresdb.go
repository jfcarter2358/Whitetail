package api

import (
	"net/http"
	"whitetail/ceresdb"
	"whitetail/datastore"
	"whitetail/utils"

	"github.com/gin-gonic/gin"
)

func Query(ctx *gin.Context) {
	var input map[string]string
	if err := ctx.BindJSON(&input); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}

	out, err := ceresdb.RawQuery(datastore.Database.Connection, datastore.Database.Auth, input["query"])
	if err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, out)
}
