package concurrency

import "time"

// 处理单个 url, 并返回 bool 结果
type WebsiteChecker func(string) bool

// 检查 url 状态
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// go 位于函数前面，开启 goroutine。
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}
	time.Sleep(2 * time.Second)
	return results
}
