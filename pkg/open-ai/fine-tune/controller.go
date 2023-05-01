package finetune

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
	fineTunes := dashboard.Group("/fine-tunes/")
	fineTunes.Use(jwt.AuthRequired(),openai.KeyRequired())
	{
		fineTunes.POST("create", h.CreateFineTune)
		fineTunes.POST("cancel/:id", h.CancelFineTune)
		fineTunes.GET("", h.GetFineTunes)
		fineTunes.GET(":id", h.GetFineTune)
		fineTunes.GET("events/:id", h.GetFineTuneEvents)
		fineTunes.DELETE(":id", h.DeleteFineTune)
	}
}	
}