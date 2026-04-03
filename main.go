package main

import (
	"fmt"
	"net/http"
	"vantagrid/handlers"
	"vantagrid/storage"
)

func main() {
	store := storage.NewStore()

	handler := &handlers.HTTPHandler{
		Store: store,
	}

	http.HandleFunc("/", handler.Handle)
	http.HandleFunc("/api/events", handler.API)

	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}
