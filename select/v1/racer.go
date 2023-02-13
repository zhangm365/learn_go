package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {

	aDuration := measureReponseTime(a)

	bDuration := measureReponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureReponseTime(url string) time.Duration {
	// 记录当前时间
	ta := time.Now()
	// http 请求 url 的内容
	http.Get(url)
	// 获取从开始时间到请求返回的时间差

	return time.Since(ta)
}
