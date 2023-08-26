package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
    "time"
	"net/http"
	"back/pkg/models"
)

func (h handler) AddUser(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
        w.WriteHeader(500)
        return
    }
    var user models.User
    json.Unmarshal(body, &user)

    user.Date = time.Now()
    queryStmt := `INSERT INTO users (id,date,name) VALUES ($1, $2, $3) RETURNING id;`
    err = h.DB.QueryRow(queryStmt, &user.Id, &user.Date, &user.Name).Scan(&user.Id)
    if err != nil {
        log.Println("failed to execute query", err)
        w.WriteHeader(500)
        return
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}