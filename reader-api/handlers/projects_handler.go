package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sieugene/jp-reader/utils"
)

func (apiCfg *ApiConfig) HandlerGetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := apiCfg.DB.GetProjects(context.Background())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get projects:%v", err))
		return
	}
	utils.RespondWithJSON(w, 201, projects)
}
