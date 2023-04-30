package prompt

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
	fineTunes := dashboard.Group("/prompt/")
	fineTunes.Use(jwt.AuthRequired(),openai.KeyRequired())
	{
		fineTunes.POST("upload", h.Uploader)
		fineTunes.POST("fine-tuner", h.FineTuner)
	}
}	
}