package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	x := Adder(2, 2)
	y := 4

	if x != y {
		t.Errorf("got '%d' but want '%d'", x, y)
	}
}

func ExampleAdder() {

	x := Adder(3, 3)
	fmt.Println(x)
	// Output: 6
}
