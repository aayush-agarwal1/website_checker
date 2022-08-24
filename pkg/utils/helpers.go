package utils

import (
	"log"
	"net/url"
	"strings"
)

func IsUrl(website string) bool {
	url, err := url.ParseRequestURI(website)
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	return strings.Contains(url.Host, ".")
}
