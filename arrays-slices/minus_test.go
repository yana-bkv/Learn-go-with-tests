package main

import "testing"

func TestMinus(t *testing.T) {
	t.Run("one numbers minus another", func(t *testing.T) {
		got := Minus(10, 5)
		want := 5

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
