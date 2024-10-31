package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sieugene/jp-reader/utils"
)

type Task struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status" enums:"processing,error,completed,waiting"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// HandlerGetTasks returns a list of tasks
// @Summary Get list of tasks
// @Description Retrieves all tasks from the database
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} Task
// @Failure 400 {string} string "Couldn't get tasks: [error message]"
// @Router /tasks [get]
func (apiCfg *ApiConfig) HandlerGetTasks(w http.ResponseWriter, r *http.Request) {
	projects, err := apiCfg.DB.GetTasks(context.Background())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get tasks:%v", err))
		return
	}
	utils.RespondWithJSON(w, 200, projects)
}
