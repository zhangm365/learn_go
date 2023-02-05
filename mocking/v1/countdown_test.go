package main

import (
	"bytes"
	"testing"
)

// 记录 Sleep() 方法调用次数
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	SpySleep := &SpySleeper{}
	// 将数据写入 buffer
	CountDown(buffer, SpySleep)

	got := buffer.String()

	// 反引号是创建 string 的另一种形式，但允许内容放置在新的一行。
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

	if SpySleep.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", SpySleep.Calls)
	}

}
