package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"
	"back/pkg/handlers"
	"back/pkg/db"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home REST API!")
	fmt.Println("Home REST API")
}

func handleRequests(DB *sql.DB) {
	h := handlers.New(DB)
	myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
	//Users
	// Create User
	myRouter.HandleFunc("/users", h.AddUser).Methods(http.MethodPost)
	// Get All Users
    myRouter.HandleFunc("/users", h.GetAllUsers).Methods(http.MethodGet)

	//Lists
	// Get All Lists
    myRouter.HandleFunc("/lists", h.GetAllLists).Methods(http.MethodGet)
	// Get Lists By UserId
	myRouter.HandleFunc("/lists/userid/{userId}", h.GetListByUserId).Methods(http.MethodGet)
	// Get List By ListId
	myRouter.HandleFunc("/lists/{listId}", h.GetList).Methods(http.MethodGet)
	// Create List
	myRouter.HandleFunc("/lists", h.AddList).Methods(http.MethodPost)
	// Update List Status By ListId
	myRouter.HandleFunc("/lists/status/{listId}", h.UpdateListStatusById).Methods(http.MethodPut)

	//Elements
	// Get All Elements
	myRouter.HandleFunc("/elements", h.GetAllElements).Methods(http.MethodGet)
	// Get Elements By ListId
	myRouter.HandleFunc("/elements/listid/{listId}", h.GetElementByListId).Methods(http.MethodGet)
	// Get ToDo Elements By ListId
	myRouter.HandleFunc("/elements-todo/listid/{listId}", h.GetElementByListIdToDo).Methods(http.MethodGet)
	// Get Doing Elements By ListId
	myRouter.HandleFunc("/elements-doing/listid/{listId}", h.GetElementByListIdDoing).Methods(http.MethodGet)
	// Get Done Elements  By ListId
	myRouter.HandleFunc("/elements-done/listid/{listId}", h.GetElementByListIdDone).Methods(http.MethodGet)
	// Create Element
	myRouter.HandleFunc("/elements", h.AddElement).Methods(http.MethodPost)
	// Update Element Status By ElementId
	myRouter.HandleFunc("/elements/status/{elementId}", h.UpdateElementStatusById).Methods(http.MethodPost)

    log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func main() {
    DB := db.Connect()
	db.CreateTableUsers(DB)
    db.CreateTableLists(DB)
	db.CreateTableElements(DB)
	fmt.Println("Server Started!")
    handleRequests(DB)
    db.CloseConnection(DB)
}

