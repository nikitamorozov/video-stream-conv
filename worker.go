package main

import (
	"encoding/json"
	"fmt"
	cfg "github.com/nikitamorozov/video-stream-conv/config/env"
	"github.com/nikitamorozov/video-stream-conv/models"
	"github.com/nikitamorozov/video-stream-conv/usecase"
	"github.com/streadway/amqp"
	"log"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()
	fmt.Println("Request tool started")
}

func main() {
	connection := config.GetString(`amqp`)
	queueName := config.GetString(`queue.name`)

	conn, err := amqp.Dial(connection)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ch, err := conn.Channel()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	forever := make(chan bool)

	uc := usecase.NewConverterUseCases()

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			job := models.Job{}
			err := json.Unmarshal(d.Body, &job)
			if err != nil {
				log.Printf(err.Error())
				return
			}

			log.Printf("source: %s dest: %s", job.Source, job.Dest)

			uc.ConvertVideo(job.Source, job.Dest)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
