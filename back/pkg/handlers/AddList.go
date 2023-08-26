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

func (h handler) AddList(w http.ResponseWriter, r *http.Request) {
        // Enable CORS by setting the appropriate headers
        w.Header().Set("Access-Control-Allow-Origin", "*")            // Allow requests from any origin
        w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") // Allow POST and OPTIONS methods
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")  // Allow Content-Type header
    
        // Handle preflight OPTIONS request (sent by browsers before the actual POST request)
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

    defer r.Body.Close()
    
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    var list models.List
    json.Unmarshal(body, &list)

    list.Id = uuid.New()
    list.Date = time.Now()
    queryStmt := `INSERT INTO lists (id,date,userId,name,status) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
    err = h.DB.QueryRow(queryStmt, &list.Id, &list.Date, &list.UserId, &list.Name, &list.Status).Scan(&list.Id)
    if err != nil {
        log.Println("failed to execute query", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}


