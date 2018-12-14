package lib

import (
	"fmt"
	"testing"
	"time"
)

func TestGetRequestId(t *testing.T) {
	for {
		fmt.Println(GenRequestId())
		time.Sleep(1 * time.Second)
	}
}
