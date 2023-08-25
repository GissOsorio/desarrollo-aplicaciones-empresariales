package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
    "time"
	"net/http"
	"back/pkg/models"
	"github.com/google/uuid"
)

func (h handler) AddElement(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
        w.WriteHeader(500)
        return
    }
    var element models.Element
    json.Unmarshal(body, &element)

    element.Id = uuid.New()
    element.Date = time.Now()
    queryStmt := `INSERT INTO elements (id,date,listId,name,status) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
    err = h.DB.QueryRow(queryStmt, &element.Id, &element.Date, &element.ListId, &element.Name, &element.Status).Scan(&element.Id)
    if err != nil {
        log.Println("failed to execute query", err)
        w.WriteHeader(500)
        return
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}