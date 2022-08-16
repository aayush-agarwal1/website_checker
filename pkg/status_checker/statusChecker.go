package status_checker

import "context"

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
}
