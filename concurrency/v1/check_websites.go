package concurrency

// 处理单个 url, 并返回 bool 结果
type WebsiteChecker func(string) bool

// 检查 url 状态
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}
