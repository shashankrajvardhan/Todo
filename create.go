package main

import (
	"encoding/json"
	"net/http"
)

func createTodo(w http.ResponseWriter, r *http.Request) {
	var a http.Header
	a = w.Header()
	a.Set("Content-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	var b *json.Encoder
	b = json.NewEncoder(w)
	b.Encode("Task added successfully")
}
