package concurrency

// 处理单个 url, 并返回 bool 结果
type WebsiteChecker func(string) bool

type result struct {
	string // url
	bool   // status
}

// 检查 url 状态
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// go 位于函数前面，开启 goroutine。
		go func(u string) {
			// send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// receive expression
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
