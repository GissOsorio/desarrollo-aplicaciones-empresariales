package handlers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"

    "PROYECTO/pkg/mocks"
    "PROYECTO/pkg/models"

    "github.com/google/uuid"
)

func AddList(w http.ResponseWriter, r *http.Request) {
    // Read to request body
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }
    var list models.List
    json.Unmarshal(body, &list)

    list.Id = (uuid.New()).String()
    
    mocks.Lists = append(mocks.Lists, list)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}