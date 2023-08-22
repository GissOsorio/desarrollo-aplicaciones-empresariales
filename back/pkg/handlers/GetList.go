package handlers

import (
	"encoding/json"
	"net/http"

	"back/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, list := range mocks.Lists {
		if list.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(list)
			break
		}
	}
}
