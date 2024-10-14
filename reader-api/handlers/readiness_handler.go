package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandlerReadiness checks server readiness
// @Summary Checks if the server is ready
// @Description Returns 200 if the server is ready
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {string} string “ok”
// @Router /healthz [get]
func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	dat, err := json.Marshal(struct{}{})
	if err != nil {
		log.Printf("Failed to marshal JSON response : %v", struct{}{})
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
