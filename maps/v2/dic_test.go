package main

import "testing"

func TestSearch(t *testing.T) {
	dic := Dic{"key1": "value1"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dic.Search("key1")
		want := "value1"
		assertStrings(t, got, want)
	})

	t.Run("uknown word", func(t *testing.T) {
		_, got := dic.Search("unknown")
		// want := "could not find the word"

		assertError(t, got, ErrNotFound)

		//assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s', but want '%s'.", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s', want '%s'.", got, want)
	}
}
