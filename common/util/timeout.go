package util

import (
	"context"
)

func DoWithTimeout(ctx context.Context, f func() error) error {
	executeChan := make(chan bool, 1)

	var handlerErr error
	go func() {
		handlerErr = f()
		// 执行完函数以后，关闭channel
		close(executeChan)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-executeChan:
	}

	return handlerErr
}
