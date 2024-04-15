package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	var a http.Header
	a = w.Header()
	a.Set("Content-Type", "application/json")
	var params map[string]string
	params = mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}
	var found bool
	found = false
	for i, item := range todos {
		if item.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			found = true
			break
		}
	}
	if found {
		var b *json.Encoder
		b = json.NewEncoder(w)
		b.Encode("Item deleted successfully")
	} else {
		var c *json.Encoder
		c = json.NewEncoder(w)
		c.Encode("Item not found")
	}
}
