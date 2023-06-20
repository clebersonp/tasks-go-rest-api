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

// Delete deletes the task in the database by its id
func Delete(w http.ResponseWriter, r *http.Request) {
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

	if rowsAffected, err := models.Delete(int64(id)); err != nil {
		log.Printf("Error trying to delete task: %v\n", err)
		payload := createPayloadError(tryAgainLaterMsg)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode payload error: %v\n", err)
		}
		return
	} else if rowsAffected == 0 {
		log.Printf("Delete Operation. Task not found for id: %d\n", id)
		msg := fmt.Sprintf("No tasks are affected for id '%d'. Make sure the task exists before deleting it.",
			id)
		payload := createPayloadError(msg)
		w.WriteHeader(http.StatusConflict)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode payload error: %v\n", err)
		}
		return
	} else if rowsAffected > 1 {
		log.Printf("More than one row was affected in the delete: '%d' rows\n", rowsAffected)
	}
	w.WriteHeader(http.StatusNoContent)
}
