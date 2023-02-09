package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	// 记录当前时间
	sa := time.Now()
	// 请求 url 的内容
	http.Get(a)
	// 获取从开始时间到请求返回的时间差
	aDuration := time.Since(sa)

	sb := time.Now()
	http.Get(b)
	bDuration := time.Since(sb)

	if aDuration < bDuration {
		return a
	}
	return b
}
