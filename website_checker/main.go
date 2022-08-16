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
	"strconv"
	"time"
)

func launchWebsiteCheckerCron(conf map[string]string) {
	concurrency, _ := strconv.Atoi(conf["concurrency"])
	delay, _ := strconv.Atoi(conf["cron.delaySeconds"])
	go func() {
		status_checker.ConcurrentStatusCheck(concurrency, delay)
	}()
	return
}

func launchWebServer(conf map[string]string) {
	host := conf["server.host"]
	port := conf["server.port"]
	writeTimeoutSeconds, _ := strconv.Atoi(conf["server.writeTimeoutSeconds"])
	readTimeoutSeconds, _ := strconv.Atoi(conf["server.readTimeoutSeconds"])
	router := mux.NewRouter()
	router.HandleFunc("/websites", api.PostWebsites).Methods(http.MethodPost)
	router.HandleFunc("/websites", api.GetWebsites).Methods(http.MethodGet)
	log.Printf("Starting server at: %s:%s", host, port)

	srv := &http.Server{
		Addr:         host + ":" + port,
		WriteTimeout: time.Second * time.Duration(writeTimeoutSeconds),
		ReadTimeout:  time.Second * time.Duration(readTimeoutSeconds),
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	return
}

func main() {
	configFile := "application.properties"
	conf := utils.ReadConfig(configFile)

	launchWebsiteCheckerCron(conf)
	launchWebServer(conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Printf("Shutting down server")
	os.Exit(0)
}
