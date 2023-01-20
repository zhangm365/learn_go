package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	got := Repeat("a", 3)
	want := "aaa"

	if got != want {
		t.Errorf("got '%s', but want '%s'", got, want)
	}
}

/*
func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}

}
*/

func ExampleRepeat() {

	str := Repeat("a", 6)
	fmt.Println(str)
	// out : "aaaaaa"
}
