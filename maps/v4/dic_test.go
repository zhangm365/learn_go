package main

import "testing"

// 字典中搜索某个 key 是否存在
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

// 在字典中添加一对 [key, value]
func TestAdd(t *testing.T) {

	t.Run("new key", func(t *testing.T) {

		dic := Dic{}
		word := "key1"
		value := "value1"
		err := dic.Add(word, value)

		assertError(t, err, nil)

		assertValue(t, dic, word, value)

	})

	t.Run("existing key", func(t *testing.T) {

		word := "key1"
		value := "value1"

		dic := Dic{word: value}

		// 对已存在的无法添加
		err := dic.Add(word, "new value1")

		assertError(t, err, ErrWordExists)

		assertValue(t, dic, word, value)

	})

}

// update the value
func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {

		word := "key1"
		value := "value1"
		dic := Dic{word: value}

		newValue := "new value1"

		err := dic.Update(word, newValue)

		assertError(t, err, nil)

		assertValue(t, dic, word, newValue)

	})

	t.Run("new word", func(t *testing.T) {

		dic := Dic{}

		word := "key1"
		value := "value1"

		err := dic.Update(word, value)

		assertError(t, err, ErrWordNotExist)

	})

}

func TestDelete(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {

		word := "key1"
		value := "value1"

		dic := Dic{word: value}

		dic.Delete(word)
		_, err := dic.Search(word)

		assertError(t, err, ErrNotFound)

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

func assertValue(t *testing.T, dic Dic, word, value string) {

	t.Helper()

	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("should find the added word:", err)
	}

	if got != value {
		t.Errorf("got '%s' want '%s'.", got, value)
	}

}
