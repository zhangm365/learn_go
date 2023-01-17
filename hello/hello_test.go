package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("zhangmao")
	want := "Hello World, chris"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
