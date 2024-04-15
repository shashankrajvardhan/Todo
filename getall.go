package main

import (
	"encoding/json"
	"net/http"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	var b http.Header
	b = w.Header()
	b.Set("Content-Type", "application/json")
	var a *json.Encoder
	a = json.NewEncoder(w)
	a.Encode(todos)
}