package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {

		// 辅助函数
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}

	}

	t.Run("saying hello to a person", func(t *testing.T) {

		got := Hello("chris", "")
		want := "Hello, chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world in Spanish", func(t *testing.T) {

		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world in French", func(t *testing.T) {

		got := Hello("Lauren", french)
		want := "Bonjour, Lauren"

		assertCorrectMessage(t, got, want)
	})

}
