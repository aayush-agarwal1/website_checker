package api

import (
	"encoding/json"
	"github.com/aayush-agarwal1/website_checker/pkg/model"
	"log"
	"net/http"
	"strings"
)

type postRequestBody struct {
	Websites []string `json:"websites"`
}

func PostWebsites(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.RequestURI)

	var websites postRequestBody

	if err := json.NewDecoder(r.Body).Decode(&websites); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, website := range websites.Websites {
		if err, wasPresent := model.InsertNewWebsite(website); err != nil {
			log.Fatalf("Encountered error while inserting website `%s` into model: %s", website, err.Error())
		} else {
			if wasPresent {
				log.Printf("Website `%s` was already in model", website)
			} else {
				log.Printf("Inserted website `%s` into model", website)
			}
		}
	}
	w.Write([]byte("Updated list of websites: " + strings.Join(model.GetWebsiteList(), ",")))
}

func GetWebsites(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	log.Printf("%s %s\n", r.Method, r.RequestURI)
	websites := strings.Split(name, ",")
	var out []byte
	out, _ = json.Marshal(model.GetWebsiteStatusMap(websites))
	w.Write(out)
}
