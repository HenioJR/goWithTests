package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {

	//table driven tests
	//are useful when you want to build a list of test cases that can be tested in the same manner
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tdt := range areaTests {
		got := tdt.shape.Area()
		if got != tdt.want {
			t.Errorf("%#v got '%v' want '%v'", tdt, got, tdt.want)
		}
	}
}
