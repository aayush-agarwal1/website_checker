package main

import (
	"github.com/aayush-agarwal1/website_checker/pkg/api"
	"github.com/aayush-agarwal1/website_checker/pkg/status_checker"
	"github.com/aayush-agarwal1/website_checker/pkg/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func launchWebsiteCheckerCron(config utils.Config) {

	concurrency := config.Cron.Concurrency
	delay := config.Cron.DelaySeconds

	go func() {
		status_checker.ConcurrentStatusCheck(concurrency, delay)
	}()
	return
}

func launchWebServer(config utils.Config) {

	host := config.Server.Host
	port := config.Server.Port
	writeTimeout := config.Server.WriteTimeoutSeconds
	readTimeout := config.Server.ReadTimeoutSeconds

	router := mux.NewRouter()
	router.HandleFunc("/websites", api.PostWebsites).Methods(http.MethodPost)
	router.HandleFunc("/websites", api.GetWebsites).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         host + ":" + port,
		WriteTimeout: time.Second * writeTimeout,
		ReadTimeout:  time.Second * readTimeout,
		Handler:      router,
	}

	log.Printf("Starting server at: %s:%s", host, port)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	return
}

func main() {
	configFile := "properties.yml"
	conf := utils.ReadConfig(configFile)

	launchWebsiteCheckerCron(conf)
	launchWebServer(conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Printf("Shutting down server")
	os.Exit(0)
}
