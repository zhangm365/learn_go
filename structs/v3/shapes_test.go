package main

import "testing"

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {

		rectangle := Rectangle{12, 6}
		got := Area(rectangle)
		want := 72.0

		if got != want {
			t.Errorf("got %.2f, but want %.2f", got, want)
		}

	})

	t.Run("circles", func(t *testing.T) {

		circle := Circle{10}
		got := Area(circle)
		want := 314.16

		if got != want {
			t.Errorf("got %.2f, but want %.2f", got, want)
		}

	})

}
