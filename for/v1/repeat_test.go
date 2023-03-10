package iteration

import "testing"

func TestIteration(t *testing.T) {

	got := Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got '%s', but want '%s'", got, want)
	}

}

// benchmark
func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a")
	}

}
