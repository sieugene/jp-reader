package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sieugene/jp-reader/utils"
)

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
	tasks, err := apiCfg.DB.GetTasks(context.Background())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get tasks:%v", err))
		return
	}
	utils.RespondWithJSON(w, 200, databaseTasksToTasks(tasks))
}

// HandlerGetTasksWithProjects returns a list of tasks with associated projects
// @Summary Get list of tasks with projects
// @Description Retrieves all tasks along with their associated projects from the database.
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} TaskWithProject
// @Failure 400 {string} string "Couldn't get tasks: [error message]"
// @Router /tasks/projects [get]
func (apiCfg *ApiConfig) HandlerGetTasksWithProjects(w http.ResponseWriter, r *http.Request) {
	tasksWithProjectsRaw, err := apiCfg.DB.GetTasksWithProjects(context.Background())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get tasks with projects:%v", err))
		return
	}
	tasksWithProjects, err := databaseTasksAndProjectsToTasksAndProjects(tasksWithProjectsRaw)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get tasks with projects:%v", err))
		return
	}

	utils.RespondWithJSON(w, 200, tasksWithProjects)
}
