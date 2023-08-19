package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "PROYECTO/pkg/mocks"
    "PROYECTO/pkg/models"
)

func GetElementByListId(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    listId := vars["listId"]

    matchingElements := []models.Element{} 

    for _, element := range mocks.Elements {
        if element.ListId == listId {
            matchingElements = append(matchingElements, element) 
        }
    }

    if len(matchingElements) == 0 {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(matchingElements)
}