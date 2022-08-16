package main

import (
	"context"
	"github.com/aayush-agarwal1/website_checker/pkg/api"
	"github.com/aayush-agarwal1/website_checker/pkg/model"
	"github.com/aayush-agarwal1/website_checker/pkg/status_checker"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func basePath(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.RequestURI)
	http.Error(w, "Invalid call, please check URI", http.StatusBadRequest)
}

func main() {
	host := "127.0.0.1"
	port := ":8080"
	router := mux.NewRouter()
	router.HandleFunc("/", basePath)
	router.HandleFunc("/websites", api.PostWebsites).Methods(http.MethodPost)
	router.HandleFunc("/websites", api.GetWebsites).Methods(http.MethodGet)
	log.Printf("Starting server at: %s%s", host, port)

	srv := &http.Server{
		Addr:         host + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("%s\n", err)
		}
	}()

	var checker status_checker.StatusChecker = status_checker.HTTPChecker{}
	go func(statusChecker status_checker.StatusChecker) {
		for {
			for _, website := range model.GetWebsiteList() {
				ctx := context.Background()
				if status, _ := statusChecker.Check(ctx, website); status {
					model.UpdateWebsiteStatus(website, model.UP)
				} else {
					model.UpdateWebsiteStatus(website, model.DOWN)
				}
			}
			time.Sleep(5 * time.Second)
		}
	}(checker)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Printf("Shutting down server")
	os.Exit(0)
}
