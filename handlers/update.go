package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/clebersonp/tasks-go-rest-api/models"
	"github.com/go-chi/chi"
)

// Update updates the task in the database
func Update(w http.ResponseWriter, r *http.Request) {
	if payload, ok := validateContentType(r.Header.Get(contentType), applicationJson); !ok {
		w.Header().Add(contentType, applicationJson)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode error payload: %v\n", err)
		}
		return
	}

	w.Header().Add(contentType, applicationJson)

	idstr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Printf("Error trying to get task's id from URLParam: %v\n", err)
		payload := createPayloadError(fmt.Sprintf("Invalid id '%s': %v", idstr, err))
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode payload error: %v\n", err)
		}
		return
	}

	var task models.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error trying to decode task's payload: %v\n", err)
		payload := createPayloadError(fmt.Sprintf("Invalid payload: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode payload error: %v\n", err)
		}
		return
	}

	if rowsAffected, err := models.Update(int64(id), task); err != nil {
		log.Printf("Error trying to update task: %v\n", err)
		payload := createPayloadError(tryAgainLaterMsg)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode payload error: %v\n", err)
		}
		return
	} else if rowsAffected == 0 {
		log.Printf("Update Operation. Task not found for id: %d\n", id)
		msg := fmt.Sprintf("No tasks are affected for id '%d'. Make sure the task exists before updating it.",
			id)
		payload := createPayloadError(msg)
		w.WriteHeader(http.StatusConflict)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode payload error: %v\n", err)
		}
		return
	} else if rowsAffected > 1 {
		log.Printf("More than one row was affected in the update: '%d' rows\n", rowsAffected)
	}
	w.WriteHeader(http.StatusOK)
}
