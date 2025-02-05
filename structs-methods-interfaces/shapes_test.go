package structs_methods_interfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("the perimeter of a rectangle with width 10 and height 10 is 40", func(t *testing.T) {
		rectangle := Rectangle{10.00, 10.00}
		got := rectangle.Perimeter()
		want := 40.0
		assertShapesCalc(t, got, want)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0}, // pode ficar mais claro explicitar assim
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}

func assertShapesCalc(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want) // t.Errorf("got %g want %g", got, want) -> um formato alternativo
	}
}
