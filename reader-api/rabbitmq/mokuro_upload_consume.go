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

	"github.com/google/uuid"
	"github.com/sieugene/jp-reader/handlers"
	"github.com/sieugene/jp-reader/internal/database"
	"github.com/sieugene/jp-reader/queue"
	"github.com/streadway/amqp"
)

func MokuroUploadConsume(apiCfq handlers.ApiConfig, config RabbitMQConfig) {
	conn, err := amqp.Dial(GetRabbitURL(config))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	for {
		msg, ok, err := ch.Get(q.Name, false)
		if err != nil {
			log.Println("Error getting message:", err)
			time.Sleep(1 * time.Second)
			continue
		}

		if ok {
			var task queue.UploadQueue
			err = json.Unmarshal(msg.Body, &task)
			if err != nil {
				log.Println("Error unmarshalling task:", err)
				if err := msg.Ack(false); err != nil {
					log.Println("Error acknowledging message:", err)
				} else {
					log.Println("Malformed task acknowledged and removed from the queue")
				}
				continue
			}

			log.Printf("Received task with ID: %s", task.ID)
			log.Printf("Processing task with ID: %s", task.ID)

			processTask(task, apiCfq)
			// TODO retry
			// if err != nil {
			// 	log.Println("Error processing task:", err)
			// 	msg.Nack(false, true)
			// 	continue
			// }

			if err := msg.Ack(false); err != nil {
				log.Println("Error acknowledging message:", err)
			} else {
				log.Printf("Task with ID: %s acknowledged", task.ID)
			}
		} else {
			log.Println("Waiting next tasks:", err)
			time.Sleep(5 * time.Second)
		}
	}
}

func updateTaskStatus(apiCfq handlers.ApiConfig, id uuid.UUID, status string) {
	err := apiCfq.DB.UpdateTaskStatus(context.Background(), database.UpdateTaskStatusParams{
		Status:    status,
		ID:        id,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		log.Println("Error while updating task:", err)
	}
}

func processTask(task queue.UploadQueue, apiCfq handlers.ApiConfig) {
	updateTaskStatus(apiCfq, task.ID, "processing")

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	files, err := os.ReadDir(task.Folder)
	if err != nil {
		log.Println("Error reading folder:", err)
		updateTaskStatus(apiCfq, task.ID, "error")
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
		updateTaskStatus(apiCfq, task.ID, "error")
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		updateTaskStatus(apiCfq, task.ID, "error")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Error response from Mokuro service:", resp.Status)
		updateTaskStatus(apiCfq, task.ID, "error")
		return
	}

	updateTaskStatus(apiCfq, task.ID, "completed")
	// Clear temp file folder
	if err := os.RemoveAll(task.Folder); err != nil {
		log.Println("Error deleting folder:", err)
	}
}
