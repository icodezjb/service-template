package util

import (
	"context"
	"errors"
	"log"
)

func DoWithTimeout(ctx context.Context, f func()) error {
	c := make(chan bool, 1)

	go func() {
		f()
		// 执行完函数以后，关闭channel
		close(c)
	}()

	select {
	case <-ctx.Done():
		log.Fatalf("[%v]redis处理超时", ctx.Value("request_id"))
		return errors.New("访问超时")
	case <-c:
		return nil
	}

	return nil
}
