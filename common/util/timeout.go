package util

import (
	"context"
	"log"

	"github.com/buchenglei/service-template/common/definition"
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
		log.Fatalf("[%v]redis处理超时", ctx.Value(definition.FieldRequestId))
		return ctx.Err()
	case <-c:
		return nil
	}

	return nil
}
