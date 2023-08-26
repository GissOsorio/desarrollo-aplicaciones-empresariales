package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"back/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) UpdateElementStatusById(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    vars := mux.Vars(r)
    id := vars["elementId"]

    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var updatedElement models.Element
    json.Unmarshal(body, &updatedElement)

    queryStmt := `UPDATE elements SET status = $2 WHERE id = $1 RETURNING id;`
    err = h.DB.QueryRow(queryStmt, &id, &updatedElement.Status).Scan(&id)
    if err != nil {
        log.Println("failed to execute query", err)
        w.WriteHeader(500)
        return
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Updated")

}