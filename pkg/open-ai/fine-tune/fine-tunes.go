package finetune

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/common/settings"
)

func (h handler) GetFineTunes(ctx *gin.Context) {
	res, err := settings.Client.Client.ListFineTunes(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error getting fine-tunes: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (h handler) GetFineTune(ctx *gin.Context) {
	fineTuneId := ctx.Param("id")
	res, err := settings.Client.Client.GetFineTune(ctx, fineTuneId)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error getting fine-tune: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (h handler) GetFineTuneEvents(ctx *gin.Context) {
	fineTuneId := ctx.Param("id")
	res, err := settings.Client.Client.ListFineTuneEvents(ctx, fineTuneId)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error getting fine-tune events: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, res)
}