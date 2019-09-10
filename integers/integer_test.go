package integers

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
