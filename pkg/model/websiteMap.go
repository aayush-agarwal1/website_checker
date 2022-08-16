package model

type State string

const (
	UP   State = "UP"
	DOWN       = "DOWN"
	INIT       = "INIT"
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

func GetWebsiteMapObject() map[string]WebsiteProperties {
	return websiteMapObject
}

func GetWebsiteStatusMap() (websiteStatusMap map[string]State) {
	websiteStatusMap = make(map[string]State)
	for website := range websiteMapObject {
		websiteStatusMap[website] = websiteMapObject[website].Status
	}
	return
}
