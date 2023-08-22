package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"back/pkg/mocks"
	"back/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateElement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedElement models.Element
	json.Unmarshal(body, &updatedElement)

	for index, element := range mocks.Elements {
		if element.Id == id {
			element.Name = updatedElement.Name
			element.Status = updatedElement.Status

			mocks.Elements[index] = element

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
