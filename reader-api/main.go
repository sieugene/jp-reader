package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sieugene/jp-reader/handlers"
	"github.com/sieugene/jp-reader/internal/database"
	"github.com/sieugene/jp-reader/rabbitmq"

	_ "github.com/lib/pq"
)

type RabbitTask struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	FileData []byte    `json:"file_data"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	portString := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	pool, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db := database.New(pool)
	apiCfq := handlers.ApiConfig{
		DB: db,
	}

	rabbitConfig := rabbitmq.RabbitMQConfig{
		User:     os.Getenv("RABBITMQ_USER"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     os.Getenv("RABBITMQ_PORT"),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	router := chi.NewRouter()

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlers.HandlerReadiness)
	v1Router.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		apiCfq.UploadHandler(w, r, rabbitmq.CreateCbSendToQueue(rabbitConfig))
	})

	v1Router.Get("/projects", apiCfq.HandlerGetProjects)
	v1Router.Post("/projects", apiCfq.HandlerCreateProjects)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	go rabbitmq.ConsumeQueue(apiCfq, rabbitConfig)

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
