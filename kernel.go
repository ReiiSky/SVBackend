package backend

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type KernelParams struct {
	Port int
}

type kernel struct {
	params    KernelParams
	ginEngine *gin.Engine
}

func NewKernel(params KernelParams) kernel {
	return kernel{
		params,
		gin.Default(),
	}
}

func (k kernel) RunServer() {
	if k.ginEngine == nil {
		panic(errors.New("gin Engine is not initialized"))
	}

	activePort := k.params.Port

	if activePort <= 1024 {
		activePort = 3000
	}

	address := fmt.Sprintf("0.0.0.0:%d", activePort)
	log.Printf("Starting server at: %s\n", address)
	k.ginEngine.Run(address)
}
