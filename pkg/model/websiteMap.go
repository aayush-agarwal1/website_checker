package model

type State string

const (
	UP             State = "UP"
	DOWN           State = "DOWN"
	INIT           State = "INIT"
	INVALID_URI    State = "INVALID_URI"
	DOES_NOT_EXIST State = "DOES_NOT_EXIST"
)

type WebsiteProperties struct {
	Status State `json:"status"`
}

var websiteMapObject map[string]WebsiteProperties

func init() {
	websiteMapObject = newWebsiteMap()
}

func newWebsiteMap() map[string]WebsiteProperties {
	return make(map[string]WebsiteProperties)
}

func InsertNewWebsite(website string) (err error, wasPresent bool) {
	err = nil
	_, wasPresent = websiteMapObject[website]
	if !wasPresent {
		websiteMapObject[website] = WebsiteProperties{
			Status: INIT,
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

func GetWebsiteMapObject(websites []string) (localWebsiteMapObject map[string]WebsiteProperties) {
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
