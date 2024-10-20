// @title JP Reader API
// @version 1.0
// @description API for JP Reader application.
// @host localhost:3000
// @BasePath /v1

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sieugene/jp-reader/handlers"
	"github.com/sieugene/jp-reader/internal/database"
	"github.com/sieugene/jp-reader/pools"
	"github.com/sieugene/jp-reader/rabbitmq"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/lib/pq"
	_ "github.com/sieugene/jp-reader/docs"
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

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlers.HandlerReadiness)
	v1Router.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		apiCfq.UploadHandler(w, r, rabbitmq.MokuroUploadTask(rabbitConfig))
	})

	v1Router.Get("/projects", apiCfq.HandlerGetProjects)

	router.Mount("/v1", v1Router)
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	go rabbitmq.MokuroUploadConsume(apiCfq, rabbitConfig)
	go pools.StartPollingProjects(apiCfq)

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
