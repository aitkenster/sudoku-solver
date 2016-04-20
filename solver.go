package main

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrIncorrectLength   = errors.New("Grid is incorrect length")
	ErrInvalidCharacters = errors.New("Grid contains invalid characters")
	ErrDuplicateValues   = errors.New("Invalid puzzle - rows/columns/sectors contain duplicate values")
)

func isValidGrid(gridInput string) error {
	if len(gridInput) != 81 {
		return ErrIncorrectLength
	}
	return validChars(gridInput)
}

func validChars(gridInput string) error {
	validChars := regexp.MustCompile(`^(\d|\s)*$`)
	if !validChars.MatchString(gridInput) {
		return ErrInvalidCharacters
	}
	return nil
}

type Grid struct {
	Squares [][]SquareVal
}

type SquareVal int

func transformIntoGrid(gridInput string) (*Grid, error) {
	g := Grid{
		Squares: [][]SquareVal{},
	}
	err := g.addRows(gridInput)
	if err != nil {
		return nil, err
	}
	err = g.checkForDuplicates()
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (g *Grid) addRows(gridInput string) error {
	for i := 0; i < len(gridInput); i += 9 {
		row := []SquareVal{}
		for j := i; j < (i + 9); j++ {
			v, err := strconv.Atoi(string(gridInput[j]))
			if err != nil {
				return err
			}
			row = append(row, SquareVal(v))
		}
		g.Squares = append(g.Squares, row)
	}
	return nil
}

type SelectionType int

const (
	ColSelection SelectionType = iota
	RowSelection
	SubMatrixSelection
)

func (g *Grid) getSelection(stype SelectionType, Location int) []SquareVal {
	selection := []SquareVal{}
	if stype == ColSelection {
		for _, s := range g.Squares {
			selection = append(selection, s[Location])
		}
	} else if stype == RowSelection {
		selection = g.Squares[Location]

	} else {
		selection = g.subMatrixes()[Location]
	}
	return selection
}

func (g *Grid) checkForDuplicates() error {
	for i := 0; i < 8; i++ {
		c := g.getSelection(ColSelection, i)
		if hasDuplicateSquares(c) {
			return ErrDuplicateValues
		}
		r := g.getSelection(RowSelection, i)
		if hasDuplicateSquares(r) {
			return ErrDuplicateValues
		}
		m := g.getSelection(SubMatrixSelection, i)
		if hasDuplicateSquares(m) {
			return ErrDuplicateValues
		}
	}
	return nil
}

func hasDuplicateSquares(squares []SquareVal) bool {
	for i, s := range squares {
		if s != 0 {
			for j := i + 1; j < len(squares); j++ {
				if s == squares[j] {
					return true
				}
			}
		}
	}
	return false
}

func (g *Grid) subMatrixes() [][]SquareVal {
	sm := make([][]SquareVal, 9)
	for i := 0; i < len(sm); i += 3 {
		for j, _ := range g.Squares {
			if j >= i && j < i+3 {
				sm[i] = append(sm[i], g.Squares[j][0:2]...)
				sm[i+1] = append(sm[i], g.Squares[j][3:5]...)
				sm[i+2] = append(sm[i], g.Squares[j][6:8]...)
			}
		}
	}
	return sm
}
