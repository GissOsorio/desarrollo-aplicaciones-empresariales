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

func AddElement(w http.ResponseWriter, r *http.Request) {
    // Read to request body
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }
    var element models.Element
    json.Unmarshal(body, &element)

    element.Id = (uuid.New()).String()
    
    mocks.Elements = append(mocks.Elements, element)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}