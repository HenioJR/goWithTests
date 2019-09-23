package mocking

import (
	"bytes"
	"testing"
)

func TestCounting(t *testing.T) {
	buffer := &bytes.Buffer{}

	Counting(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
