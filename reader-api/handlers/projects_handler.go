package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sieugene/jp-reader/utils"
)

// Project represents a project structure.
// @Description Represents a project with its details.
// @ID Project
// @Property ID string `json:"id"` // The unique identifier for the project.
// @Property createdAt string `json:"createdAt"` // The timestamp of project creation in ISO 8601 format.
// @Property updatedAt string `json:"updatedAt"` // The timestamp of the last update in ISO 8601 format.
// @Property Name string `json:"name"` // The name of the project.
// @Property Images array[string] `json:"images"` // A list of images associated with the project.
// @Property OcrData object `json:"ocrData"` // OCR data associated with the project.
type Project struct {
	ID        uuid.UUID   `json:"ID"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Name      string      `json:"Name"`
	Images    []string    `json:"Images"`
	OcrData   interface{} `json:"OcrData"`
}

// HandlerGetProjects returns a list of projects
// @Summary Get list of projects
// @Description Retrieves all projects from the database
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {array} Project
// @Failure 400 {string} string "Couldn't get projects: [error message]"
// @Router /projects [get]
func (apiCfg *ApiConfig) HandlerGetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := apiCfg.DB.GetProjects(context.Background())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get projects:%v", err))
		return
	}
	utils.RespondWithJSON(w, 201, projects)
}
