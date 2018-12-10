package message

import "context"

type Module struct {
}

func New() *Module {
	return &Module{}
}

// AsyncSendUserLoginMessage 异步发送消息模块
// Notic 该方法为异步执行的方法，执行后立刻返回，由后台goroutine继续执行
// 该类方法必须使用Async开头表明这已经是一个异步方法了，无需外部再通过go关键字开启新的goroutine
// 同时该异步方法必须需要正确处理ctx的超时
func (*Module) AsyncSendUserLoginMessage(ctx context.Context, account string) {
	go func() {
		// send message
		// log .......
		select {
		case <-ctx.Done():
			// Error 超时
		}

	}()
}
