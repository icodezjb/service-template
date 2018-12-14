package lib

import (
	"github.com/rs/xid"
)

var (
	requestIdChan chan string
)

func init() {
	requestIdChan = make(chan string, 300)

	go generator()
}

func generator() {
	for {
		select {
		case requestIdChan <- makeRequestId():
		}
	}

}

func makeRequestId() string {
	return xid.New().String()
}

func GenRequestId() string {
	return <-requestIdChan
}
