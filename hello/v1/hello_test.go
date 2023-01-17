package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("chris")
	want := "Hello World, chris"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
