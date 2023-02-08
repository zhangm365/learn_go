package main

import (
	"fmt"
	"net/http"
)

// 检查 url 是否合法
func CheckWebsite(url string) bool {
	response, err := http.Head(url)
	if err != nil {
		return false
	}

	// 200 status code
	if response.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func main() {
	fmt.Println(CheckWebsite("https://www.baidu.com/"))
}
