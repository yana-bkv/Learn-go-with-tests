package iteration

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	got := Matches("foo", "seafood")
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %s, got %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
