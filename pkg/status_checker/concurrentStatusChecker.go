package status_checker

import (
	"context"
	"github.com/aayush-agarwal1/website_checker/pkg/model"
	"time"
)

var checker StatusChecker

func ConcurrentStatusCheck(concurrency int) {
	checker = HTTPChecker{}

	channel := make(chan string, concurrency*2)

	for i := 0; i < concurrency; i++ {
		go func(channel chan string) {
			for {
				select {
				case website := <-channel:
					ctx := context.Background()
					if status, _ := checker.Check(ctx, website); status {
						model.UpdateWebsiteStatus(website, model.UP)
					} else {
						model.UpdateWebsiteStatus(website, model.DOWN)
					}
				}
			}
		}(channel)
	}

	for {
		for _, website := range model.GetWebsiteList() {
			channel <- website
		}
		time.Sleep(30 * time.Second)
	}
}
