package main

import "testing"

func TestAdder(t *testing.T) {
	sum := Adder(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
