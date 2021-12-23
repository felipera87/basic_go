package surfaces_test

import (
	"math"
	. "sub-tests/surfaces"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 12}
		expectedArea := float64(120)
		receivedArea := rectangle.Area()

		if expectedArea != receivedArea {
			// Fatalf is same as Errorf, except that it will stop executing the other errors
			t.Fatalf("Expected %f but got %f", expectedArea, receivedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		expectedArea := float64(math.Pi * 100)
		receivedArea := circle.Area()

		if expectedArea != receivedArea {
			t.Fatalf("Expected %f but got %f", expectedArea, receivedArea)
		}
	})
}
