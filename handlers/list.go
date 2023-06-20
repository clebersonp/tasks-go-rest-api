package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/clebersonp/tasks-go-rest-api/models"
)

// List gets all tasks from database
func List(w http.ResponseWriter, r *http.Request) {

	w.Header().Add(contentType, applicationJson)

	tasks, err := models.GetAll()
	if err != nil {
		log.Printf("Error trying to get_all tasks from database: %v\n", err)
		payload := createPayloadError(tryAgainLaterMsg)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode payload error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Printf("Error trying to encode success payload: %v\n", err)
	}
}
