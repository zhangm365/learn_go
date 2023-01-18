package main

import "testing"

func TestHello(t *testing.T) {

	// 可以在函数中声明另一个函数并将它们分配给变量
	assertCorrectMessage := func(t *testing.T, got, want string) {

		// 辅助函数
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("chris")
		want := "Hello, chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {

		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

}

/* version 2
// 将断言重构为函数
func assertCorrectMessage(t *testing.T, got, want string) {

	// 辅助函数
	t.Helper()

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}

}
*/
