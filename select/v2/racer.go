package racer

import "net/http"

func Racer(a, b string) (winner string, err error) {
	select { // 同时在多个 channel 上等待返回信息，第一个发送值的 channel 「胜出」，case 代码执行
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
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
