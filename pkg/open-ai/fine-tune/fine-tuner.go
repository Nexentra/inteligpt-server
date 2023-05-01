package finetune

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/common/settings"
	"github.com/nexentra/inteligpt/pkg/httputil"
	openai "github.com/sashabaranov/go-openai"
)

func (h handler) CreateFineTune(ctx *gin.Context) {
	body := openai.FineTuneRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	response, err := settings.Client.Client.CreateFineTune(ctx, body)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error creating fine-tune: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h handler) CancelFineTune(ctx *gin.Context) {
	fineTuneId := ctx.Param("id")
	res, err := settings.Client.Client.CancelFineTune(ctx, fineTuneId)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error canceling fine-tune: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (h handler) DeleteFineTune(ctx *gin.Context) {
	fineTuneId := ctx.Param("id")
	res, err := settings.Client.Client.DeleteFineTune(ctx, fineTuneId)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error deleting fine-tune: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, res)
}