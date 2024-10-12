package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sieugene/jp-reader/internal/database"
	"github.com/sieugene/jp-reader/utils"
)

func (apiCfq *ApiConfig) HandlerCreateProjects(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}
	if len(params.Name) == 0 || len(params.Link) == 0 {
		utils.RespondWithError(w, 400, fmt.Sprintln("Name and Link is required"))
		return
	}

	project, err := apiCfq.DB.CreateProject(context.Background(), database.CreateProjectParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      params.Name,
		Link:      params.Link,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
	utils.RespondWithJSON(w, 201, project)
}

func (apiCfg *ApiConfig) HandlerGetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := apiCfg.DB.GetProjects(context.Background())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get projects:%v", err))
		return
	}
	utils.RespondWithJSON(w, 201, projects)
}
