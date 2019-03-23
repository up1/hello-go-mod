package world

import "testing"

func TestHello(t *testing.T) {
	want := "World."
	if got := World(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
