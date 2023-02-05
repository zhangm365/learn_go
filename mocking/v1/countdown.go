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
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

// func CountDown(out *bytes.Buffer) {
/*
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
*/

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countStart; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	// 1s
	// time.Sleep(1 * time.Second)
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	// 数据写入标准输出
	CountDown(os.Stdout, sleeper)
}
