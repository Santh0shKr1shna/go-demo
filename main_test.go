package shape

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestNewReactangleWithValidValues(t *testing.T) {
	t.Run("Testing the new rectangle with valid values", func(t *testing.T) {
		// arrange
		height := 10
		breadth := 12

		want := Rectangle {
			height: height,
			breadth: breadth,
		}

		// action
		got, err :=  NewRectangle(height, breadth)

		// errors mitigation
		if err != nil {
			t.Errorf("ERR: %s", err.Error())
		}

		// assertion
		if got.breadth != want.breadth || got.height != want.height {
			t.Errorf("Value mismatch")
		}
	})
}

func TestNewReactangleWithInvalidValues(t *testing.T) {

	t.Run("Testing the new rectangle with invalid values (negative)", func(t *testing.T) {
		// arrange
		height := -10
		breath := -12

		// action
		_, err :=  NewRectangle(height, breath)

		// assertion 1
		if assert.Nil(t, err) {
			t.Errorf("Negative value not detected")
		}

		// assertion 2
		if assert.NotEqual(t, "Rectangle cannot have negative or zero dimensions", err.Error()) {
			t.Errorf("Different error printed\nERR: %s", err.Error())
		}
	})
}

func TestRectangleAreaWithValidValues(t *testing.T) {
	t.Run("Test rectangle area function", func(t *testing.T) {
		// arrange
		h := 10
		b := 12
		want := b * h

		// action
		rect, _ := NewRectangle(h, b)
		got := rect.CalculateArea()

		// assertion
		if want != got {
			t.Errorf("got: %d want: %d", got, want)
		}
	})
}

func TestRectanglePerimeterWithValidValues(t *testing.T) {
	t.Run("Testing rectangle perimeter function", func(t *testing.T) {
		// arrange
		h := 5
		b := 6
		want := 2 * (b + h)

		// action
		rect, _ := NewRectangle(h, b)
		got := rect.CalculatePerimeter()

		// assertion
		if want != got {
			t.Errorf("got: %d\nwant:= %d", got, want)
		}
	})
}

// SQUARE TESTS

func TestSquareWithValidValues(t *testing.T) {
	// arrange
	side := 4

	// action
	got, err :=  NewSquare(side)

	// errors mitigation
	if assert.NotNil(t, err) {
		t.Errorf("ERR: %s", err.Error())
	} else if assert.Nil(t, got) {
		t.Error("Nil object returned")
	}
}

func TestSquareWithInvalidValues(t *testing.T) {
	// arrange
	side := -4

	// action
	got, err :=  NewSquare(side)

	// assertion
	if assert.Nil(t, err) {
		t.Error("No error thrown by constructor")
	} else if assert.NotNil(t, got) {
		t.Error("Non nil object thrown by constructor")
	}
}

func TestGetSideMethodOfStructSquare(t *testing.T) {
	side := 10

	square, _ := NewSquare(side)

	if assert.Equal(t, side, square.GetSide()) {
		t.Error("Values mismatch")
	}
}

func TestSquareAreaWithValidValues(t *testing.T) {
	side := 4
	want := side * side

	sq, _ := NewSquare(side)
	got := sq.CalculateArea()

	if assert.Equal(t, want, got) {
		t.Errorf("Values mismatch\nGot: %d\tWant: %d", got, want)
	}
}

func TestSquarePerimeterWithValidValues(t *testing.T) {
	side := 5
	want := 4 * side

	sq, _ := NewSquare(side)
	got := sq.CalculatePerimeter()

	if assert.Equal(t, want, got) {
		t.Errorf("Values mismatch\nGot: %d\tWant: %d", got, want)
	}
}