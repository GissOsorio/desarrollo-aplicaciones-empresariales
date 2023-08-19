package handlers

import (
	"encoding/json"
	"net/http"

	"desarrollo-aplicaciones-empresariales/pkg/mocks"
)

func GetAllLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Lists)
}
