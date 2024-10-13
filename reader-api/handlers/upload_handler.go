package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sieugene/jp-reader/internal/database"
	"github.com/sieugene/jp-reader/queue"
)

type UploadResponse struct {
	Message string `json:"message"`
}

type SendToQueueFunc func(task queue.UploadQueue) error

func (apiConfig *ApiConfig) UploadHandler(w http.ResponseWriter, r *http.Request, sendToQueue SendToQueueFunc) {
	r.ParseMultipartForm(10 << 20)

	paramTitle := r.FormValue("title")
	if paramTitle == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["file"]
	if len(files) == 0 {
		http.Error(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	folderPath := fmt.Sprintf("temps/%s/", paramTitle)
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		http.Error(w, "Error creating upload folder", http.StatusInternalServerError)
		return
	}

	for _, handler := range files {
		file, err := handler.Open()
		if err != nil {
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		outFile, err := os.Create(folderPath + handler.Filename)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		if _, err = io.Copy(outFile, file); err != nil {
			http.Error(w, "Error copying file", http.StatusInternalServerError)
			return
		}
	}

	dbWaitingTask, err := apiConfig.DB.CreateTask(context.Background(), database.CreateTaskParams{
		ID:        uuid.New(),
		Title:     paramTitle,
		Status:    "waiting",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		http.Error(w, "Error creating task record", http.StatusInternalServerError)
		return
	}

	task := queue.UploadQueue{
		ID:     dbWaitingTask.ID,
		Title:  paramTitle,
		Folder: folderPath,
	}

	err = sendToQueue(task)
	if err != nil {
		http.Error(w, "Error sending task to queue", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Task has been queued"))
}
