package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sieugene/jp-reader/internal/database"
)

type Project struct {
	ID        uuid.UUID   `json:"id"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Name      string      `json:"name"`
	Images    []string    `json:"images"`
	OcrData   interface{} `json:"ocrData"`
}

func databaseProjectToProject(dbProject database.Project) Project {
	return Project{
		ID:        dbProject.ID,
		CreatedAt: dbProject.CreatedAt,
		UpdatedAt: dbProject.UpdateAt,
		Name:      dbProject.Name,
		Images:    dbProject.Images,
		OcrData:   dbProject.OcrData,
	}
}

func databaseProjectsToProjects(dbProjects []database.Project) []Project {
	projects := []Project{}
	for _, dbProject := range dbProjects {
		projects = append(projects, databaseProjectToProject(dbProject))
	}
	return projects
}

type Task struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status" enums:"processing,error,completed,waiting"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func databaseTaskToTask(dbTask database.Task) Task {
	return Task{
		ID:        dbTask.ID,
		Title:     dbTask.Title,
		Status:    dbTask.Status,
		CreatedAt: dbTask.CreatedAt,
		UpdatedAt: dbTask.UpdatedAt,
	}
}

func databaseTasksToTasks(dbTasks []database.Task) []Task {
	tasks := []Task{}
	for _, dbTask := range dbTasks {
		tasks = append(tasks, databaseTaskToTask(dbTask))
	}
	return tasks
}

type TaskWithProject struct {
	Task
	Project Project `json:"project"`
}

func databaseTaskAndProjectToTaskAndProject(dbTaskWithProject database.GetTasksWithProjectsRow) (TaskWithProject, error) {
	task := databaseTaskToTask(database.Task{
		ID:        dbTaskWithProject.ID,
		Title:     dbTaskWithProject.Title,
		Status:    dbTaskWithProject.Status,
		CreatedAt: dbTaskWithProject.CreatedAt,
		UpdatedAt: dbTaskWithProject.UpdatedAt,
	})
	var dbProject database.Project
	err := json.Unmarshal(dbTaskWithProject.Project, &dbProject)
	if err != nil {
		return TaskWithProject{}, fmt.Errorf("failed to unmarshal project data: %w", err)
	}

	project := databaseProjectToProject(dbProject)

	return TaskWithProject{
		Task:    task,
		Project: project,
	}, nil
}

func databaseTasksAndProjectsToTasksAndProjects(dbTasksWithProjects []database.GetTasksWithProjectsRow) ([]TaskWithProject, error) {
	tasksAndProjects := []TaskWithProject{}
	for _, taskWithProject := range dbTasksWithProjects {
		var data, err = databaseTaskAndProjectToTaskAndProject(taskWithProject)
		if err != nil {
			return []TaskWithProject{}, err
		}
		tasksAndProjects = append(tasksAndProjects, data)
	}
	return tasksAndProjects, nil
}
