package hello

import "testing"

func TestWelcome(t *testing.T) {
	r := Welcome()
	if r != "Hello" {
		t.Errorf("Welcome() return %s; but expected 'Hello'", r)
	}
}
