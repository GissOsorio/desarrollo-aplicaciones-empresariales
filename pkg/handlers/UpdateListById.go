package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"desarrollo-aplicaciones-empresariales/pkg/mocks"
	"desarrollo-aplicaciones-empresariales/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedList models.List
	json.Unmarshal(body, &updatedList)

	for index, list := range mocks.Lists {
		if list.Id == id {
			list.Name = updatedList.Name
			list.Status = updatedList.Status

			mocks.Lists[index] = list

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
