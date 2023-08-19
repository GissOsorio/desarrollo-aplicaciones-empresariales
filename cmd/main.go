package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "PROYECTO/pkg/handlers"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Article REST API!")
    fmt.Println("Article REST API")
}

func handleRequests() {
    // create a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/lists", handlers.GetAllLists).Methods(http.MethodGet)
    myRouter.HandleFunc("/lists/{id}", handlers.GetList).Methods(http.MethodGet)
    myRouter.HandleFunc("/lists", handlers.AddList).Methods(http.MethodPost)
    log.Fatal(http.ListenAndServe(":8074", myRouter))
}

func main() {
    handleRequests()
}