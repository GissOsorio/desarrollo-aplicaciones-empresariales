package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "PROYECTO/pkg/mocks"
)

func GetElement(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for _, element := range mocks.Elements {
        if element.Id == id {
            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(element)
            break
        }
    }
}