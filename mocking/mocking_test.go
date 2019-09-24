package main

import (
	"bytes"
	"testing"
)

//used to mock
type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperSpy := &SleeperSpy{}

	Countdown(buffer, sleeperSpy)

	got := buffer.String()
	//create a string with a single quote allows line break
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if sleeperSpy.Calls != 4 {
		t.Errorf("sleeperSpy.Calls: got '%d' want 4", sleeperSpy.Calls)
	}
}
