package main

import (
	"fmt"
	"items-api/internal/handlers"
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/items", handlers.GetItems)

	fmt.Println("server started :8080")
	http.ListenAndServe(":8080", router)
}
