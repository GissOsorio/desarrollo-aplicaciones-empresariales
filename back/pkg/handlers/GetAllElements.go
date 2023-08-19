package handlers

import (
	"encoding/json"
	"net/http"

	"back/pkg/mocks"
)

func GetAllElements(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Elements)
}
