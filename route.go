package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func route() {
	var router *mux.Router
	router = mux.NewRouter()
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todo", createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")
	router.HandleFunc("/todos", deleteTodos).Methods("DELETE").Queries("id", "{id}")
	fmt.Println("Server running on port 8000")

	var a error
	a = http.ListenAndServe(":8000", router)
	log.Fatal(a)
}
