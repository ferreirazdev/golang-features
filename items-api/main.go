package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/items", GetItems)
	http.ListenAndServe(":8090", nil)
}

func GetItems(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "no-items\n")
}
