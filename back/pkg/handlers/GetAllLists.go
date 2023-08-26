package handlers

import (
	"encoding/json"
	"net/http"
	"log"

	"back/pkg/models"
)

func (h handler) GetAllLists(w http.ResponseWriter, r *http.Request) { 
    enableCors(&w)
    results, err := h.DB.Query("SELECT * FROM lists;")
    if err != nil {
        log.Println("failed to execute query", err)
        w.WriteHeader(500)
        return
    }

    var lists = make([]models.List, 0)
    for results.Next() {
        var list models.List
        err = results.Scan(&list.Id, &list.Date, &list.UserId, &list.Name, &list.Status)
        if err != nil {
            log.Println("failed to scan", err)
            w.WriteHeader(500)
            return
        }

        lists = append(lists, list)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(lists)	

}