package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func basePath(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	http.Error(w, "Invalid call, please check URI", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", basePath)
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
