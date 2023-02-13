package main

import (
	"log"

	"net/http"

	"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"
	"github.com/itokun99/ms-go-boilerplate/core/config/db"
	"github.com/itokun99/ms-go-boilerplate/core/handler"
	"github.com/itokun99/ms-go-boilerplate/core/middleware"
	"github.com/itokun99/ms-go-boilerplate/core/model"
	"github.com/itokun99/ms-go-boilerplate/jsonxml/controller"
	"github.com/itokun99/ms-go-boilerplate/jsonxml/repository"
	"github.com/itokun99/ms-go-boilerplate/jsonxml/usecase"
	"go.elastic.co/apm/module/apmgin/v2"

	"github.com/itokun99/ms-go-boilerplate/core/config/env"
	log2 "github.com/itokun99/ms-go-boilerplate/core/config/log"
)

func main() {

	// init router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	env.NewEnv(".env")

	// init amp gin
	router.Use(apmgin.Middleware(router))

	// setup no route
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"responseCode": "4048800", "responseMessage": "Page not found"})
	})

	// setup contract
	if err := errcntrct.InitContract(env.Config.JSONPathFile); err != nil {
		log.Println("main : init contract", err)
		panic(err)
	}

	dbBase := db.NewDB(env.Config).DB
	err := dbBase.Debug().Migrator().AutoMigrate(model.Logs{})
	if err != nil {
		panic(err)
	}

	logCustom := log2.NewLogCustom(env.Config, dbBase)

	// repository
	jsonXmlRepo := repository.NewJsonXmlRepository(dbBase, logCustom)

	// usecase
	jsonXmlUsecase := usecase.NewJsonXmlUseCase(jsonXmlRepo, logCustom, env.Config)

	// router group
	parserRouter := router.Group("/parser")

	// error handler
	errorHandler := handler.NewErrorHandler()

	// middleware
	middleware.NewErrorMiddleware(parserRouter, errorHandler, logCustom)

	// controller
	controller.NewJsonXmlController(parserRouter, logCustom, jsonXmlUsecase, env.Config)

	if err := router.Run(env.Config.Host + ":" + env.Config.Port); err != nil {
		log.Fatal("main : error starting server", err)
	}
}
