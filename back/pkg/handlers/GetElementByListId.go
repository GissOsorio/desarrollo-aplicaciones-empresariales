package handlers

import (
	"encoding/json"
	"net/http"

	"desarrollo-aplicaciones-empresariales/pkg/mocks"
	"desarrollo-aplicaciones-empresariales/pkg/models"
	"github.com/gorilla/mux"
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
