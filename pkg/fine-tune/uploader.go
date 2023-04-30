package finetune

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/httputil"
	openai "github.com/sashabaranov/go-openai"

	"github.com/nexentra/inteligpt/pkg/common/settings"
)

func (h handler) Uploader(ctx *gin.Context) {
	req := openai.FileRequest{
		FileName: "my_file",
		FilePath: "./predict.jsonl",
		Purpose:  "fine-tune",
	}
	res, err := settings.Client.Client.CreateFile(context.Background(), req)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
