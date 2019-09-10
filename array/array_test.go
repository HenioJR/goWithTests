package array

import "testing"

func TestAdd(t *testing.T) {
	numbers := []int{1, 2, 3, 4}

	got := Add(numbers)
	want := 10

	if got != want {
		t.Errorf("got '%d' want '%d' array '%v'", got, want, numbers)
	}
}
