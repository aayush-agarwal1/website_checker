package status_checker

import (
	"context"
	"log"
	"net/http"
)

type HTTPChecker struct {
}

func (checker HTTPChecker) Check(ctx context.Context, name string) (status bool, err error) {
	if resp, err := http.Get("https://" + name); err == nil {
		if resp.StatusCode < 299 {
			log.Printf("website %s is up\n", name)
			return true, nil
		} else {
			log.Printf("website %s is down with status code %d\n", name, resp.StatusCode)
			return false, nil
		}
	} else {
		log.Printf(err.Error())
		return false, err
	}
}
