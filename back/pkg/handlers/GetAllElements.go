package handlers

import (
	"encoding/json"
	"net/http"
	"log"

	"back/pkg/models"
)

func (h handler) GetAllElements(w http.ResponseWriter, r *http.Request) { 
    enableCors(&w)
    results, err := h.DB.Query("SELECT * FROM elements;")
    if err != nil {
        log.Println("failed to execute query", err)
        w.WriteHeader(500)
        return
    }

    var elements = make([]models.Element, 0)
    for results.Next() {
        var element models.Element
        err = results.Scan(&element.Id, &element.Date, &element.ListId, &element.Name, &element.Status)
        if err != nil {
            log.Println("failed to scan", err)
            w.WriteHeader(500)
            return
        }

        elements = append(elements, element)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(elements)	

}