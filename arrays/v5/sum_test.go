package main

import (
	"reflect"
	"testing"
)

func TestSumTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("test the empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		checkSums(t, got, want)
	})

}
