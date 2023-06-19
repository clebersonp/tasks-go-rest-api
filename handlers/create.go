package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/clebersonp/tasks-go-rest-api/models"
)

// Create creates given task in database
func Create(w http.ResponseWriter, r *http.Request) {
	if payload, ok := validateContentType(r.Header.Get(contentType), applicationJson); !ok {
		w.Header().Add(contentType, applicationJson)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode error payload: %v\n", err)
		}
		return
	}

	w.Header().Add(contentType, applicationJson)
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task);
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error trying to decode body payload: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newTask, err := models.Insert(task)
	if err != nil {
		payload := createPayloadError(fmt.Sprintf("Something wrong happens trying to insert tasks: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode error payload: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		log.Printf("Error trying to encode success payload: %v\n", err)
	}
}
