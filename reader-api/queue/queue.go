package queue

import "github.com/google/uuid"

type UploadQueue struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Folder string    `json:"folder"`
}
