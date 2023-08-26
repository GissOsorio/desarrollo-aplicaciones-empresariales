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
    (*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}