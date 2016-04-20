package main

import "testing"

func TestValidGrid(t *testing.T) {
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
			"103456789103456789123006709123456789123456789123456789123456789123456789123456789",
			nil,
		},
	}

	for _, test := range tests {
		actual := Valid(test.grid)
		if actual != test.expected {
			t.Errorf("expected %b for grid %#v, but got %b",
				test.expected, test.grid, actual)
		}
	}
}

func TestTransformIntoGrid(t *testing.T) {
	gridInput := "102004070000902800009003004000240006000107000400068000200800700007501000080400109"
	valGrid, err := transformIntoGrid(gridInput)
	if err != nil {
		t.Fatalf("expected valid grid, got %#v", err)
	}

	if len(valGrid.Cells) != len(gridInput)/9 {
		t.Errorf("expected grid to contain %#v squares, instead got %#v", len(gridInput)/9, len(valGrid.Cells))
	}

	expected := 0

	if valGrid.Cells[0][4] != expected {
		t.Errorf("expected square to equal %#v, instead got %#v", expected, valGrid.Cells[0][4])
	}

	invalidGrids := []string{
		"123456789123456789123456789123456789123456789123456789123456789123456789123456789", // duplicate values in column
		"112004070000902800009003004000240006000107000400068000200800700007501000080400109", //duplicate values in row
		"102004070010902800009003004000240006000107000400068000200800700007501000080400109", //duplicate values in sector
	}

	for _, input := range invalidGrids {
		_, err = transformIntoGrid(input)
		if err == nil {
			t.Fatalf("expected %#v, got %#v", ErrDuplicateValues, nil)
		}
	}
}

func TestSimpleSolve(t *testing.T) {
	gridInput := "000008070060003090092000058400760000001004000000200000100807000003090000500006400"
	valGrid, err := transformIntoGrid(gridInput)
	if err != nil {
		t.Fatalf("expected valid grid, got:", err)
	}
	expected := "315928674864573192792641358438769521251384967976215843149857236683492715527136489"

	actual, err := valGrid.simpleSolve()
	if err != nil {
		t.Fatalf("expected solution, got:", err)
	}

	if actual != expected {
		t.Fatalf("expected %#v, got %#v", expected, actual)
	}

}
