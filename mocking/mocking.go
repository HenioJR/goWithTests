package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

//used to real implementation
type SleeperReal struct{}

func (s *SleeperReal) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprint(w, "Go!")
}

func main() {
	sleeper := &SleeperReal{}
	Countdown(os.Stdout, sleeper)
}
