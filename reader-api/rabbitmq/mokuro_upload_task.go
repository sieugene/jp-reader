package rabbitmq

import (
	"encoding/json"

	"github.com/sieugene/jp-reader/queue"
	"github.com/streadway/amqp"
)

func MokuroUploadTask(config RabbitMQConfig) func(task queue.UploadQueue) error {
	return func(task queue.UploadQueue) error {

		conn, err := amqp.Dial(GetRabbitURL(config))
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			return err
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			queue.UPLOAD_QUEUE_KEY,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}

		body, err := json.Marshal(task)
		if err != nil {
			return err
		}

		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		return err
	}

}
