package main

import (
	"fmt"
	// "os"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/auth"
	"github.com/nexentra/inteligpt/pkg/common/db"
	filecontrol "github.com/nexentra/inteligpt/pkg/open-ai/file-control"
	finetune "github.com/nexentra/inteligpt/pkg/open-ai/fine-tune"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() error {
	fmt.Println("starting up api service")

	return nil
}

func setupRouter() *gin.Engine {

	fmt.Println("Go-Gin REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	// port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	// dsn := fmt.Sprintf(
	// 	"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USERNAME"),
	// 	os.Getenv("DB_TABLE"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("SSL_MODE"),
	// )

	// fmt.Println(dsn)

	r := gin.Default()
	h := db.InitDatabase(dbUrl)
	d := r.Group("/dashboard")

	auth.RegisterRoutes(r,h)
	finetune.RegisterRoutes(d,r,h)
	filecontrol.RegisterRoutes(d,r,h)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to inteligpt!!",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func main() {
	r := setupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
