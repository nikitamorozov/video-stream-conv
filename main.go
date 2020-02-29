package main

import (
	"fmt"
	"github.com/labstack/echo"
	cfg "github.com/nikitamorozov/video-stream-conv/config/env"
	"github.com/nikitamorozov/video-stream-conv/config/middleware"
	"github.com/nikitamorozov/video-stream-conv/delivery/http"
	"github.com/nikitamorozov/video-stream-conv/repository"
	"github.com/nikitamorozov/video-stream-conv/usecase"
	"github.com/streadway/amqp"
)

var configMain cfg.Config

func init() {
	configMain = cfg.NewViperConfig()
	fmt.Println("Request tool started")
}

func main() {
	domain := configMain.GetString(`server.domain`)
	postfix := configMain.GetString(`converter.postfix`)
	connection := configMain.GetString(`amqp`)
	queueName := configMain.GetString(`queue.name`)

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	conn, err := amqp.Dial(connection)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	repo := repository.NewQueueRepository(conn)
	uc := usecase.NewQueueUseCases(repo)

	http.NewManagerHttpHandler(e, queueName, uc, domain, postfix)

	_ = e.Start(configMain.GetString("server.address"))
}
