package filecontrol

import (
	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/middlware/jwt"
	openai "github.com/nexentra/inteligpt/middlware/open-ai"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(dashboard *gin.RouterGroup,router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
{
	fineTunes := dashboard.Group("/files/")
	fineTunes.Use(jwt.AuthRequired(),openai.KeyRequired())
	{
		fineTunes.POST("upload", h.CreateFile)
		fineTunes.DELETE(":id", h.DeleteFile)
		fineTunes.GET("", h.GetFiles)
		fineTunes.GET(":id", h.GetFile)
	}
}	
}