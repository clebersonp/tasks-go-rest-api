package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/clebersonp/tasks-go-rest-api/models"
	"github.com/go-chi/chi"
)

// Get gets a task by id
func Get(w http.ResponseWriter, r *http.Request) {

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

	task, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Error trying to get task by id from database: %v\n", err)
		if sql.ErrNoRows.Error() == err.Error() {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		payload := createPayloadError(tryAgainLaterMsg)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error trying to encode payload error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Printf("Error trying to encode success payload: %v\n", err)
	}
}
