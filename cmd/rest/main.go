package main

import (
	"log"
	"os"

	backend "github.com/ReiiSky/SharingVisionBackend"
	"github.com/ReiiSky/SharingVisionBackend/application"
	"github.com/ReiiSky/SharingVisionBackend/infrastructure"
	"github.com/ReiiSky/SharingVisionBackend/infrastructure/databases"
	"github.com/ReiiSky/SharingVisionBackend/packages"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("error loading .env file")
	}

	articleRepository := setupArticleRepository()
	validate := validator.New(validator.WithRequiredStructEnabled())

	kernel := backend.NewKernel(backend.KernelProps{
		Port:              GetEnvInt(3000, "PORT"),
		ArticleRepository: articleRepository,
		Validate:          validate,
	})

	setupController(kernel)

	kernel.RunServer()
}

func setupArticleRepository() packages.ArticleRepository {
	databaseConnection, err := databases.NewDatabaseConnection(
		databases.DatabaseConnectionConfig{
			Hostname: os.Getenv("MYSQL_HOSTNAME"),
			Port:     GetEnvInt(3306, "MYSQL_PORT"),
			DBname:   os.Getenv("MYSQL_DBNAME"),
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Queries:  os.Getenv("MYSQL_QUERIES"),
		},
	)

	if err != nil {
		log.Fatalf("error on initializing database with message: %s", err.Error())
	}

	return infrastructure.NewArticleRepository(databaseConnection)
}

func setupController(k backend.Kernel) {
	createArticleController := application.NewCreateArticleController(k)
	getArticleController := application.NewGetArticleController(k)
	getDetailArticleController := application.NewGetDetailArticleController(k)
	updateArticleController := application.NewUpdateArticleController(k)

	k.RegisterHandlerFn(func(e *gin.Engine) {
		e.POST("/article", createArticleController.Handle)
		e.GET("/article", getArticleController.Handle)
		e.GET("/article/:id", getDetailArticleController.Handle)
		e.PUT("/article/:id", updateArticleController.Handle)
	})
}
