package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {

		nums := []int{1, 2, 3, 4, 5}
		got := Sum(nums)
		want := 15

		if got != want {
			t.Errorf("got %d, but want %d. '%v'", got, want, nums)
		}

	})

}
