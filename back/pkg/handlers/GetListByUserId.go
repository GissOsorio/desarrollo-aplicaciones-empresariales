package handlers

import (
	"encoding/json"
	"net/http"
	"log"

	"back/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetListByUserId(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    vars := mux.Vars(r)
    userId := vars["userId"]

    queryStmt := `SELECT * FROM lists WHERE userid = $1 ;`
    results, err := h.DB.Query(queryStmt, userId)
    if err != nil {
        log.Println("failed to execute query", err)
        w.WriteHeader(500)
        return
    }
	defer results.Close()

	var lists []models.List 

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
    
    err = json.NewEncoder(w).Encode(lists)
    if err != nil {
        log.Println("failed to encode JSON", err)
        w.WriteHeader(500)
        return
    }
}