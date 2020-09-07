package hadlers

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func respondJson(w http.ResponseWriter, status int, payload interface{})  {
	response, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJson(w, code, map[string]string{"error": message})
}

func getIdentifier(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])

	if err != nil {
		return id, err
	}

	return id, nil
}
