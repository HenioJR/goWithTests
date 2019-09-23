package injection

import (
	"fmt"
	"io"
)

func Salute(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}
