package finetune

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/httputil"
	openai "github.com/sashabaranov/go-openai"

	"github.com/nexentra/inteligpt/pkg/common/settings"
)

func (h handler) Uploader(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	req := openai.FileRequest{
		FilePath: filename,
		Purpose:  "fine-tune",
	}
	res, err := settings.Client.Client.CreateFile(context.Background(), req)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
