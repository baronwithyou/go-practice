package main

import "testing"

// TestAddDigits ...
func TestAddDigits(t *testing.T) {
	got := addDigits(123)

	want := int64(6)
	if got != want {
		t.Fatalf("want: %d, got: %d", want, got)
	}
}
