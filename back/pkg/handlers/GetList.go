package handlers

import (
	"encoding/json"
	"net/http"
	"log"

	"back/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetList(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["listId"]

    queryStmt := `SELECT * FROM lists WHERE id = $1 ;`
    results, err := h.DB.Query(queryStmt, id)
    if err != nil {
        log.Println("failed to execute query", err)
        w.WriteHeader(500)
        return
    }

    var list models.List
    for results.Next() {
        err = results.Scan(&list.Id, &list.Date, &list.UserId, &list.Name, &list.Status)
        if err != nil {
            log.Println("failed to scan", err)
            w.WriteHeader(500)
            return
        }
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(list)
}