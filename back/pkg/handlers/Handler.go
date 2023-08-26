package handlers

import (
    "database/sql"
    "net/http"
)

type handler struct {
    DB *sql.DB
}

func New(db *sql.DB) handler {
    return handler{db}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}