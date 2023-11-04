package routes

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khozaei/healthio/ent"
)

var router = gin.Default()

type HandlerConfig struct {
	Client  *ent.Client
	Context context.Context
}

var handlerConfig HandlerConfig

func Run() {
	listenAddress := ":3000"
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	} else {
		listenAddress = os.Getenv("LISTEN_ADDRESS")
	}
	setRoutes()
	router.Run(listenAddress)
}

func SetHandlerConfig(conf HandlerConfig) {
	handlerConfig.Client = conf.Client
	handlerConfig.Context = conf.Context
}

func setRoutes() {
	v1 := router.Group("v1")
	setPatientRoutes(v1)
	setReceptionRoutes(v1)
	setVisitsRoutes(v1)
}
