package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	select { // 同时在多个 channel 上等待返回信息，第一个发送值的 channel 「胜出」，对应的 case 代码执行
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	// 使用 select 时，time.After() 是一个很好的函数。
	case <-time.After(10 * time.Second): // 超时时间是 10s。
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
