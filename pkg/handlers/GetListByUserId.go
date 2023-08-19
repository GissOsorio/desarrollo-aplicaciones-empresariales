package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "PROYECTO/pkg/mocks"
    "PROYECTO/pkg/models"
)

func GetListByUserId(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["userId"]

    matchingLists := []models.List{} 

    for _, list := range mocks.Lists {
        if list.UserId == userId {
            matchingLists = append(matchingLists, list) 
        }
    }

    if len(matchingLists) == 0 {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(matchingLists)
}