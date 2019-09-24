package injection

import (
	"fmt"
	"io"
	"net/http"
)

func Salute(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func HandlerSalute(rw http.ResponseWriter, r *http.Request) {
	Salute(rw, "world")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerSalute))

	if err != nil {
		fmt.Println(err)
	}
}
