package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dic{"test": "this is a just test map"}
	got := dictionary.Search("test")
	want := "this is a just test map"
	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s', but want '%s'.", got, want)
	}

}
