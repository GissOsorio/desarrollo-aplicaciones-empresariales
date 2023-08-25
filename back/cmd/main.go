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
	// create a new instance of a mux router
	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage)
	// Get All Lists
	// myRouter.HandleFunc("/lists", handlers.GetAllLists).Methods(http.MethodGet)
	// Get List By ListId
	// myRouter.HandleFunc("/lists/{id}", handlers.GetList).Methods(http.MethodGet)
	// Get List By UserId
	// myRouter.HandleFunc("/lists/userid/{userId}", handlers.GetListByUserId).Methods(http.MethodGet)
	// Add List
	// myRouter.HandleFunc("/lists", handlers.AddList).Methods(http.MethodPost)
	// Update List By ListId
	// myRouter.HandleFunc("/lists/{id}", handlers.UpdateList).Methods(http.MethodPut)
	// Get All Elements
	// myRouter.HandleFunc("/elements", handlers.GetAllElements).Methods(http.MethodGet)
	// Get Elements By ElementId
	// myRouter.HandleFunc("/elements/{id}", handlers.GetElement).Methods(http.MethodGet)
	// Get Elements By ListId
	// myRouter.HandleFunc("/elements/listid/{listId}", handlers.GetElementByListId).Methods(http.MethodGet)
	// Add Element
	// myRouter.HandleFunc("/elements", handlers.AddElement).Methods(http.MethodPost)
	// Update Element By ElementId
	// myRouter.HandleFunc("/elements/{id}", handlers.UpdateElement).Methods(http.MethodPut)
	
	myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
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
	myRouter.HandleFunc("/elements/status/{elementId}", h.UpdateElementStatusById).Methods(http.MethodPut)
	


	//myRouter.HandleFunc("/elements/{id}", h.GetArticle).Methods(http.MethodGet)
    //myRouter.HandleFunc("/elements", h.AddArticle).Methods(http.MethodPost)
    //myRouter.HandleFunc("/elements/{id}", h.UpdateArticle).Methods(http.MethodPut)
    
    log.Fatal(http.ListenAndServe(":8028", myRouter))

}

func main() {
    DB := db.Connect()
    db.CreateTableLists(DB)
	db.CreateTableElements(DB)
    handleRequests(DB)
    db.CloseConnection(DB)
}
