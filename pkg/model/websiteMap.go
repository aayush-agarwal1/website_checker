package model

import (
	"github.com/aayush-agarwal1/website_checker/pkg/utils"
)

type State string

const (
	UP             State = "UP"
	DOWN           State = "DOWN"
	INIT           State = "INIT"
	INVALID_URL    State = "INVALID_URL"
	DOES_NOT_EXIST State = "DOES_NOT_EXIST"
)

type WebsiteProperties struct {
	Status State `json:"status"`
}

var websiteMapObject map[string]WebsiteProperties

func init() {
	websiteMapObject = newWebsiteMap()
}

//Singleton Pattern
func newWebsiteMap() map[string]WebsiteProperties {
	if websiteMapObject == nil {
		return make(map[string]WebsiteProperties)
	} else {
		return nil
	}
}

func InsertNewWebsite(website string) (wasPresent bool) {
	_, wasPresent = websiteMapObject[website]
	if !wasPresent {
		if isURL := utils.IsUrl("https://" + website); isURL {
			websiteMapObject[website] = WebsiteProperties{
				Status: INIT,
			}
		} else {
			websiteMapObject[website] = WebsiteProperties{
				Status: INVALID_URL,
			}
		}
	}
	return
}

func GetWebsiteList() (websites []string) {
	for website := range websiteMapObject {
		websites = append(websites, website)
	}
	return
}

func GetValidWebsiteList() (websites []string) {
	for website, properties := range websiteMapObject {
		if INVALID_URL != properties.Status {
			websites = append(websites, website)
		}
	}
	return
}

func GetWebsitePropertiesMap(websites []string) (localWebsiteMapObject map[string]WebsiteProperties) {
	if (len(websites) == 0) || (len(websites) == 1 && websites[0] == "") {
		localWebsiteMapObject = websiteMapObject
	} else {
		localWebsiteMapObject = make(map[string]WebsiteProperties)
		for _, website := range websites {
			websitePropertyMapping, isPresent := websiteMapObject[website]
			if isPresent {
				localWebsiteMapObject[website] = websitePropertyMapping
			} else {
				localWebsiteMapObject[website] = WebsiteProperties{Status: DOES_NOT_EXIST}
			}
		}
	}
	return
}

func GetWebsiteStatusMap(websites []string) (websiteStatusMap map[string]string) {
	websiteStatusMap = make(map[string]string)
	if (len(websites) == 0) || (len(websites) == 1 && websites[0] == "") {
		for website := range websiteMapObject {
			websiteStatusMap[website] = string(websiteMapObject[website].Status)
		}
	} else {
		for _, website := range websites {
			websiteProperties, isPresent := websiteMapObject[website]
			if isPresent {
				websiteStatusMap[website] = string(websiteProperties.Status)
			} else {
				websiteStatusMap[website] = string(DOES_NOT_EXIST)
			}
		}
	}
	return
}

func UpdateWebsiteStatus(website string, status State) {
	if websiteProperties, isPresent := websiteMapObject[website]; isPresent {
		websiteProperties.Status = status
		websiteMapObject[website] = websiteProperties
	}
}
