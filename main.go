package main

import (
	"fmt"
	"github.com/nikitamorozov/video-stream-conv/delivery/http"
	"github.com/nikitamorozov/video-stream-conv/usecase"
	//"github.com/go-redis/redis"
	"github.com/labstack/echo"
	cfg "github.com/nikitamorozov/video-stream-conv/config/env"
	"github.com/nikitamorozov/video-stream-conv/config/middleware"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()
	fmt.Println("Request tool started")
}

func main() {
	domain := config.GetString(`server.domain`)
	postfix := config.GetString(`converter.postfix`)
	//redisAddress := config.GetString(`redis.address`)
	//redisPassword := config.GetString(`redis.password`)
	//redisDb := config.GetInt(`redis.db`)
	//
	//redisClient := redis.NewClient(&redis.Options{
	//	Addr:     redisAddress,
	//	Password: redisPassword,
	//	DB:       redisDb,
	//})

	//redisClient.Get().

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	// Use cases layer
	useCase := usecase.NewConverterUseCases()

	// Delivery layer
	http.NewConverterHttpHandler(e, useCase, domain, postfix)

	_ = e.Start(config.GetString("server.address"))
}
