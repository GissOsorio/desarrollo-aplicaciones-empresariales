package handlers

import (
	"encoding/json"
	"net/http"
	"log"

	"back/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetElementByListIdDoing(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    listId := vars["listId"]

    queryStmt := `SELECT * FROM elements WHERE listid = $1 AND status = 'doing';`
    results, err := h.DB.Query(queryStmt, listId)
    if err != nil {
        log.Println("failed to execute query", err)
        w.WriteHeader(500)
        return
    }
	defer results.Close()

	var elements []models.Element 

    for results.Next() {
        var element models.Element
        err = results.Scan(&element.Id, &element.Date,&element.ListId, &element.Name, &element.Status)
        if err != nil {
            log.Println("failed to scan", err)
            w.WriteHeader(500)
            return
        }
        elements = append(elements, element) 
    }
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    
    err = json.NewEncoder(w).Encode(elements)
    if err != nil {
        log.Println("failed to encode JSON", err)
        w.WriteHeader(500)
        return
    }
}