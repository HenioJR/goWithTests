package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Countdown(w io.Writer) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(w, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
