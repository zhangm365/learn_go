package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{3, 6})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}

}
