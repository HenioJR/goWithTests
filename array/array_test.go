package array

import "testing"

func TestAdd(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Add(numbers)
	want := 15

	if got != want {
		t.Errorf("got '%d' want '%d' array '%v'", got, want, numbers)
	}
}
