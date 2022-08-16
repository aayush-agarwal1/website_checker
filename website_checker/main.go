package main

import (
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func basePath(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.RequestURI)
	http.Error(w, "Invalid call, please check URI", http.StatusBadRequest)
}

func main() {
	var host string = ""
	var port string = ":8080"
	router := mux.NewRouter()
	router.HandleFunc("/", basePath)
	log.Printf("Starting Server at: %s%s", host, port)
	err := http.ListenAndServe(host+port, router)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
