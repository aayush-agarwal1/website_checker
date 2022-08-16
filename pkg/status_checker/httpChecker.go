package status_checker

import (
	"context"
	"log"
	"net/http"
)

type HTTPChecker struct {
}

func (checker HTTPChecker) Check(ctx context.Context, name string) (status bool, err error) {
	workerNo := ctx.Value("workerNo")
	if resp, err := http.Get("https://" + name); err == nil {
		if resp.StatusCode < 299 {
			log.Printf("Website `%s` is up with status code `%d`, workerNo: `%d`\n", name, resp.StatusCode, workerNo)
			return true, nil
		} else {
			log.Printf("Website `%s` is down with status code `%d`, workerNo: `%d`\n", name, resp.StatusCode, workerNo)
			return false, nil
		}
	} else {
		log.Printf("Error: `%s`, workerNo: `%d`\n", err.Error(), workerNo)
		return false, err
	}
}
