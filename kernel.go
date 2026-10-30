package backend

import (
	"errors"
	"fmt"
	"log"

	"github.com/ReiiSky/SharingVisionBackend/packages"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	MIN_PORT     = 1024
	DEFAULT_PORT = 3000
)

type KernelProps struct {
	Port              int
	ArticleRepository packages.ArticleRepository
	Validate          *validator.Validate
}

type Kernel struct {
	props  KernelProps
	engine *gin.Engine
}

func NewKernel(props KernelProps) Kernel {
	engine := gin.Default()
	engine.Use(cors.Default())

	return Kernel{
		props,
		engine,
	}
}

func (k Kernel) RegisterHandlerFn(fn func(*gin.Engine)) {
	fn(k.engine)
}

func (k Kernel) Validator() *validator.Validate {
	return k.props.Validate
}

func (k Kernel) ArticleRepository() packages.ArticleRepository {
	return k.props.ArticleRepository
}

func (k Kernel) RunServer() {
	if k.engine == nil {
		panic(errors.New("gin Engine is not initialized"))
	}

	activePort := k.props.Port

	if activePort <= MIN_PORT {
		activePort = DEFAULT_PORT
	}

	address := fmt.Sprintf("0.0.0.0:%d", activePort)
	log.Printf("Starting server at: %s\n", address)

	if err := k.engine.Run(address); err != nil {
		log.Fatalf("engine stopped with error: %s", err.Error())
	}
}
