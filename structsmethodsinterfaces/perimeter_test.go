package structsmethodsinterfaces_test

import (
	"testing"

	. "example.com/go-test/structsmethodsinterfaces"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("%#v got %g but want %g", shape, got, want)
		}
	}

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "rectangle",
			shape: Rectangle{Width: 12.0, Height: 6.0},
			want:  72.0,
		},
		{
			name:  "circle",
			shape: Circle{Radius: 10},
			want:  314.1592653589793,
		},
		{
			name:  "triangle",
			shape: Triangle{Base: 12, Height: 6},
			want:  36.0,
		},
	}

	for _, tc := range areaTests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			checkArea(t, tc.shape, got, tc.want)
		})
	}
}
