package hello

import (
	"testing"

	"example.com/hello/world"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHelloWorld(t *testing.T) {
	want := "World."
	if got := world.World(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
