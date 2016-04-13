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
			t.Errorf("expected %b for grid %s, but got %b",
				test.expected, test.grid, actual)
		}
	}
}
