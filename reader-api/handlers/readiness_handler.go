package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

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
