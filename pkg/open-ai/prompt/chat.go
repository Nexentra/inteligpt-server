package prompt

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/common/settings"
	"github.com/nexentra/inteligpt/pkg/httputil"
	openai "github.com/sashabaranov/go-openai"
)

func (h handler) Chat(ctx *gin.Context) {
	body := openai.FineTuneRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	response, err := settings.Client.Client.CreateFineTune(context.Background(), body)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error creating fine-tune: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
