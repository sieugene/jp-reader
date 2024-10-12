package rabbitmq

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/sieugene/jp-reader/handlers"
	"github.com/sieugene/jp-reader/internal/database"
	"github.com/sieugene/jp-reader/queue"
	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	User     string
	Password string
	Host     string
	Port     string
}

func GetRabbitURL(config RabbitMQConfig) string {
	return "amqp://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/"
}

func ConsumeQueue(apiCfq handlers.ApiConfig, config RabbitMQConfig) {
	conn, err := amqp.Dial(GetRabbitURL(config))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"upload_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, ok, err := ch.Get(q.Name, true)
		if err != nil {
			log.Println("Error getting message:", err)
			continue
		}

		if ok {
			var task queue.UploadQueue
			err = json.Unmarshal(msg.Body, &task)
			if err != nil {
				log.Println("Error unmarshalling task:", err)
				continue
			}

			processTask(task, apiCfq)
		}
	}
}

func processTask(task queue.UploadQueue, apiCfq handlers.ApiConfig) {
	apiCfq.DB.UpdateTaskStatus(context.Background(), database.UpdateTaskStatusParams{
		Status:    "processing",
		ID:        task.ID,
		UpdatedAt: time.Now().UTC(),
	})

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	files, err := os.ReadDir(task.Folder)
	if err != nil {
		log.Println("Error reading folder:", err)
		apiCfq.DB.UpdateTaskStatus(context.Background(), database.UpdateTaskStatusParams{
			ID:        task.ID,
			Status:    "error",
			UpdatedAt: time.Now().UTC(),
		})
		return
	}

	for _, file := range files {
		filePath := filepath.Join(task.Folder, file.Name())
		fileData, err := os.Open(filePath)
		if err != nil {
			log.Println("Error opening file:", err)
			continue
		}
		defer fileData.Close()

		part, err := writer.CreateFormFile("file", file.Name())
		if err != nil {
			log.Println("Error creating form file:", err)
			continue
		}

		if _, err = io.Copy(part, fileData); err != nil {
			log.Println("Error copying file data:", err)
			continue
		}
	}

	if err := writer.Close(); err != nil {
		log.Println("Error closing writer:", err)
		return
	}

	requestEndpoint := os.Getenv("MOKURO_SERVICE") + "/upload/" + task.Title
	req, err := http.NewRequest("POST", requestEndpoint, &buf)
	if err != nil {
		log.Println("Error creating request:", err)
		apiCfq.DB.UpdateTaskStatus(context.Background(), database.UpdateTaskStatusParams{
			ID:        task.ID,
			Status:    "error",
			UpdatedAt: time.Now().UTC(),
		})
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		apiCfq.DB.UpdateTaskStatus(context.Background(), database.UpdateTaskStatusParams{
			ID:        task.ID,
			Status:    "error",
			UpdatedAt: time.Now().UTC(),
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Error response from Mokuro service:", resp.Status)
		apiCfq.DB.UpdateTaskStatus(context.Background(), database.UpdateTaskStatusParams{
			ID:        task.ID,
			Status:    "error",
			UpdatedAt: time.Now().UTC(),
		})
		return
	}

	apiCfq.DB.UpdateTaskStatus(context.Background(), database.UpdateTaskStatusParams{
		ID:        task.ID,
		Status:    "completed",
		UpdatedAt: time.Now().UTC(),
	})

	if err := os.RemoveAll(task.Folder); err != nil {
		log.Println("Error deleting folder:", err)
	}
}

func CreateCbSendToQueue(config RabbitMQConfig) func(task queue.UploadQueue) error {
	return func(task queue.UploadQueue) error {
		rabbitConfig := RabbitMQConfig{
			User:     os.Getenv("RABBITMQ_USER"),
			Password: os.Getenv("RABBITMQ_PASSWORD"),
			Host:     os.Getenv("RABBITMQ_HOST"),
			Port:     os.Getenv("RABBITMQ_PORT"),
		}

		conn, err := amqp.Dial(GetRabbitURL(rabbitConfig))
		if err != nil {
			return err
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			return err
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"upload_queue",
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
