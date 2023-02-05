package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countStart = 3

type Sleeper interface {
	Sleep()
}

// 默认休眠器
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)

}

// func CountDown(out *bytes.Buffer) {
func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countStart; i > 0; i-- {
		// time.Sleep(1 * time.Second)
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	// 1s
	// time.Sleep(1 * time.Second)
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	// 数据写入标准输出
	CountDown(os.Stdout, sleeper)
}
