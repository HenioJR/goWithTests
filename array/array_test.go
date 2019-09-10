package array

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	numbers := []int{1, 2, 3, 4}

	got := Add(numbers)
	want := 10

	if got != want {
		t.Errorf("got '%d' want '%d' array '%v'", got, want, numbers)
	}
}

func TestSumSlices(t *testing.T) {
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	got := SumSlices(sliceA, sliceB)
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func TestSumWithoutFirst(t *testing.T) {
	sliceA := []int{1, 2, 3}

	checkTest := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}

	t.Run("sum some slices", func(t *testing.T) {
		sliceB := []int{4, 5, 6}

		got := SumWithoutFirst(sliceA, sliceB)
		want := []int{5, 11}

		checkTest(t, got, want)

	})

	t.Run("sum empty slice", func(t *testing.T) {
		sliceC := []int{}
		got := SumWithoutFirst(sliceA, sliceC)
		want := []int{5, 0}

		checkTest(t, got, want)
	})

}
