package main

import "testing"

func TestIsValidGrid(t *testing.T) {
	tests := []struct {
		grid     string
		expected error
	}{
		{
			"1",
			ErrIncorrectLength,
		}, {

			"1a34567891xxxx6789123456789123456789123456789123456789123456789123456789123456789",
			ErrInvalidCharacters,
		}, {
			"123456789123456789123456789123456789123456789123456789123456789123456789123456789",
			nil,
		}, {
			"1 34567891 3456789123  67 9123456789123456789123456789123456789123456789123456789",
			nil,
		},
	}

	for _, test := range tests {
		actual := isValidGrid(test.grid)
		if actual != test.expected {
			t.Errorf("expected %b for grid %#v, but got %b",
				test.expected, test.grid, actual)
		}
	}
}

func TestTransformIntoGrid(t *testing.T) {
	gridInput := "123456789123456789123456789123456789123456789123456789123456789123456789123456789"
	grid, err := transformIntoGrid(gridInput)
	if err != nil {
		t.Fatalf("expected valid grid, got %#v", err)
	}

	if len(grid.Squares) != len(gridInput) {
		t.Errorf("expected grid to contain %#v squares, instead got %#v", len(gridInput), len(grid.Squares))
	}

	expected := Square{
		Row:    1,
		Column: 1,
		Value:  2,
	}

	if grid.Squares[10] != expected {
		t.Errorf("expected square to equal %#v, instead got %#v", expected, grid.Squares[10])
	}

}
