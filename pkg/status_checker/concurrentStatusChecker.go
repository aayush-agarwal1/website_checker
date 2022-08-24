package status_checker

import (
	"context"
	"github.com/aayush-agarwal1/website_checker/pkg/model"
	"time"
)

func worker(channel chan string, workerNo int, checker StatusChecker) {
	for {
		select {
		case website := <-channel:
			ctx := context.WithValue(context.Background(), "workerNo", workerNo)
			if status, _ := checker.Check(ctx, website); status {
				model.UpdateWebsiteStatus(website, model.UP)
			} else {
				model.UpdateWebsiteStatus(website, model.DOWN)
			}
		}
	}
}

func ConcurrentStatusCheck(concurrency int, delay time.Duration) {

	channel := make(chan string, concurrency*2)
	var checker StatusChecker = HTTPChecker{}

	for i := 1; i <= concurrency; i++ {
		go worker(channel, i, checker)
	}

	for {
		for _, website := range model.GetValidWebsiteList() {
			channel <- website
		}
		time.Sleep(time.Second * delay)
	}
}
