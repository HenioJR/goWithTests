package injection

import (
	"bytes"
	"testing"
)

func TestSalute(t *testing.T) {

	//type buffer of bytes package implements the interface Writer
	buffer := bytes.Buffer{}

	Salute(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
