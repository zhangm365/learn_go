package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

/*
	type SpySleeper struct {
		Calls int
	}

	func (s *SpySleeper) Sleep() {
		s.Calls++
	}
*/

func TestCountDown(t *testing.T) {
	t.Run("print 3 to Go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		// SpySleep := &SpySleeper{}
		// 将数据写入 buffer
		// CountDown(buffer, SpySleep)
		CountDown(buffer, &CountdownOperationSpy{})
		got := buffer.String()

		// 反引号是创建 string 的另一种形式，但允许内容放置在新的一行。
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

		/*
			if SpySleep.Calls != 4 {
				t.Errorf("not enough calls to sleeper, want 4 got %d", SpySleep.Calls)
			}
		*/
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationSpy{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durSlept)
	}
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type SpyTime struct {
	durSlept time.Duration
}

func (s *SpyTime) Sleep(du time.Duration) {
	s.durSlept = du
}
