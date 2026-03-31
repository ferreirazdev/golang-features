package handlers

import (
	"fmt"
	"net/http"
)

func GetItems(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "no-items\n")
}
