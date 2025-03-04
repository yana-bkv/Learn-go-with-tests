package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is the test"}

	t.Run("search for key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is the test"

		assertStrings(t, got, want)
	})

	t.Run("no key found", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("not existing key", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test2"
		value := "second test"

		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})
	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "this is the test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	key := "test"
	value := "this is the test"

	t.Run("existing key", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		newValue := "new test"

		dictionary.Update(key, newValue)

		assertDefinition(t, dictionary, key, newValue)
	})
	t.Run("not existing key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update(key, value)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{key: "test definition"}

		err := dictionary.Delete(key)

		assertError(t, err, nil)

		_, err = dictionary.Search(key)
		assertError(t, err, ErrNotFound)
	})
	t.Run("not existing key", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{}
		err := dictionary.Delete(key)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal(err)
	}
	assertStrings(t, got, value)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
