package filecontrol

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/httputil"
	openai "github.com/sashabaranov/go-openai"

	"github.com/nexentra/inteligpt/pkg/common/settings"
)

func (h handler) CreateFile(ctx *gin.Context) {
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
	res, err := settings.Client.Client.CreateFile(ctx, req)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}


func (h handler) DeleteFile(ctx *gin.Context) {
	fineTuneId := ctx.Param("id")
	err := settings.Client.Client.DeleteFile(ctx, fineTuneId)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error deleting file: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Successfully deleted file",
	})
}

func (h handler) GetFile(ctx *gin.Context) {
	fineTuneId := ctx.Param("id")
	res,err := settings.Client.Client.GetFile(ctx, fineTuneId)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error getting file: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (h handler) GetFiles(ctx *gin.Context) {
	res,err := settings.Client.Client.ListFiles(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error getting files: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, res)
}