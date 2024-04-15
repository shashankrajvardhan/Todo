package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

func deleteTodos(w http.ResponseWriter, r *http.Request) {
	var a http.Header
	a = w.Header()
	a.Set("Content-Type", "application/json")
	var queryParams url.Values
	queryParams = r.URL.Query()
	ids, ok := queryParams["id"]
	if !ok {
		http.Error(w, "No IDs provided", http.StatusBadRequest)
		return
	}
	var idList []int
	idList = []int{}
	for _, idStr := range ids {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid todo ID", http.StatusBadRequest)
			return
		}
		idList = append(idList, id)
	}
	var deleted int
	deleted = 0
	for _, id := range idList {
		for i, item := range todos {
			if item.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				deleted++
				break
			}
		}
	}
	if deleted > 0 {
		var b *json.Encoder
		b = json.NewEncoder(w)
		b.Encode("Items deleted successfully")
	} else {
		var c *json.Encoder
		c = json.NewEncoder(w)
		c.Encode("No items found to delete")
	}
}
