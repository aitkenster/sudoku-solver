package main

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrIncorrectLength   = errors.New("Grid is incorrect length")
	ErrInvalidCharacters = errors.New("Grid contains invalid characters")
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
	Squares []Square
}

type Square struct {
	Row    int
	Column int
	Value  int
}

func transformIntoGrid(gridInput string) (*Grid, error) {
	g := Grid{
		Squares: []Square{},
	}
	g.addSquares(gridInput)
	g.addRowsAndColumns()
	return &g, nil
}

func (g *Grid) addSquares(gridInput string) error {
	for _, val := range gridInput {
		v, err := strconv.Atoi(string(val))
		if err != nil {
			return ErrInvalidCharacters
		}
		g.Squares = append(g.Squares, Square{
			Value: v,
		})
	}
	return nil
}

func (g *Grid) addRowsAndColumns() {
	rowNum := 0
	for i := 0; i < len(g.Squares); i += 9 {
		colNum := 0
		for j := i; j < (i + 8); j++ {
			g.addRow(j, rowNum)
			g.addColumn(j, colNum)
			colNum++
		}
		rowNum++
	}
	return
}

func (g *Grid) addRow(pos, rowNum int) {
	g.Squares[pos].Row = rowNum
}

func (g *Grid) addColumn(pos, colNum int) {
	g.Squares[pos].Column = colNum
}
