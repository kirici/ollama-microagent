package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(2, 6)
	want := 12

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
