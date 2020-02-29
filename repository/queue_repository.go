package repository

import (
	"encoding/json"
	"github.com/nikitamorozov/video-stream-conv/models"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type myQueueRepo struct {
	conn *amqp.Connection
}

func NewQueueRepository(connection *amqp.Connection) QueueRepository {
	return &myQueueRepo{
		conn: connection,
	}
}

func (m *myQueueRepo) Queue(name string, job models.Job) error {
	ch, err := m.conn.Channel()

	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	content, err := json.Marshal(job)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(content),
		})

	if err != nil {
		return err
	}

	return nil
}
