package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{2, 3}
	got := Area(rectangle)
	want := 6.0
	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}

}
