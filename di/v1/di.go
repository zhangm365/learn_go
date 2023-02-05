package main

import (
	"fmt"
	"io"
	"net/http"
)

// func Greet(writer *bytes.Buffer, name string) {
func Greet(writer io.Writer, name string) {
	// io.Writer: general usage

	// fmt.Printf("Hello, %s", name) // 标准输出

	// writer 把 string 参数发送到 bytes.Buffer
	fmt.Fprintf(writer, "Hello, %s", name)

}

// 互联网
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
