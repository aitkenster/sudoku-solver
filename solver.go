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
	ErrNoSolution        = errors.New("Unsolvable puzzle")
	ErrReachedLimit      = errors.New("Square already has a value of 9")
)

func Valid(gridInput string) error {
	if len(gridInput) != 81 {
		return ErrIncorrectLength
	}
	return validChars(gridInput)
}

func validChars(gridInput string) error {
	validChars := regexp.MustCompile(`^(\d)*$`)
	if !validChars.MatchString(gridInput) {
		return ErrInvalidCharacters
	}
	return nil
}

type Grid struct {
	Cells [][]int
}

func transformIntoGrid(gridInput string) (*Grid, error) {
	g := Grid{
		Cells: [][]int{},
	}
	err := g.addRows(gridInput)
	if err != nil {
		return nil, err
	}
	if g.hasDuplicates() {
		return nil, ErrDuplicateValues
	}
	return &g, nil
}

func (g *Grid) addRows(gridInput string) error {
	for i := 0; i < len(gridInput); i += 9 {
		row := []int{}
		for j := i; j < (i + 9); j++ {
			v, err := strconv.Atoi(string(gridInput[j]))
			if err != nil {
				return err
			}
			row = append(row, v)
		}
		g.Cells = append(g.Cells, row)
	}
	return nil
}

type SelectionType int

const (
	ColSelection SelectionType = iota
	RowSelection
	SubMatrixSelection
)

func (g *Grid) getSelection(stype SelectionType, Location int) []int {
	selection := []int{}
	if stype == ColSelection {
		for _, s := range g.Cells {
			selection = append(selection, s[Location])
		}
	} else if stype == RowSelection {
		selection = g.Cells[Location]

	} else {
		selection = g.subMatrixes()[Location]
	}
	return selection
}

func (g *Grid) hasDuplicates() bool {
	selectionTypes := []SelectionType{ColSelection, RowSelection, SubMatrixSelection}
	for i, _ := range g.Cells {
		for _, st := range selectionTypes {
			s := g.getSelection(st, i)
			if hasDuplicateCells(s) {
				return true
			}
		}
	}
	return false
}

func hasDuplicateCells(cells []int) bool {
	for i, s := range cells {
		if s != 0 {
			for j := i + 1; j < len(cells); j++ {
				if s == cells[j] {
					return true
				}
			}
		}
	}
	return false
}

func (g *Grid) subMatrixes() [][]int {
	sm := make([][]int, 9)
	for i := 0; i < len(sm); i += 3 {
		for j, _ := range g.Cells {
			if j >= i && j < i+3 {
				sm[i] = append(sm[i], g.Cells[j][0:3]...)
				sm[i+1] = append(sm[i+1], g.Cells[j][3:6]...)
				sm[i+2] = append(sm[i+2], g.Cells[j][6:9]...)
			}
		}
	}
	return sm
}

type UnsolvedCell struct {
	Row    int
	Column int
}

func (g *Grid) simpleSolve() (string, error) {
	unsolved := g.Unsolved()
	err := g.findValidNumbers(unsolved, 0)
	if err != nil {
		return "", err
	}
	return g.toString(), nil
}

func (g *Grid) Unsolved() []UnsolvedCell {
	unsolved := []UnsolvedCell{}
	for i, row := range g.Cells {
		for j, s := range row {
			if s == 0 {
				u := UnsolvedCell{
					Row:    i,
					Column: j,
				}
				unsolved = append(unsolved, u)
			}
		}
	}
	return unsolved
}

func (g *Grid) findValidNumbers(ul []UnsolvedCell, current int) error {
	if g.isSolved() {
		return nil
	}
	cell := ul[current]
	if err := g.increase(cell.Row, cell.Column); err != nil {
		if current == 0 {
			return ErrNoSolution
		}
		g.reset(cell.Row, cell.Column)
		return g.findValidNumbers(ul, current-1)
	}
	if current != (len(ul) - 1) {
		return g.findValidNumbers(ul, current+1)
	}
	return nil
}

func (g *Grid) increase(r, c int) error {
	if g.Cells[r][c] == 9 {
		return ErrReachedLimit
	}
	g.Cells[r][c]++
	if g.hasDuplicates() {
		return g.increase(r, c)
	}
	return nil
}

func (g *Grid) reset(r, c int) {
	g.Cells[r][c] = 0
}

func (g *Grid) isSolved() bool {
	if len(g.Unsolved()) == 0 {
		if !g.hasDuplicates() {
			return true
		}
	}
	return false
}

func (g *Grid) toString() string {
	var s string
	for _, r := range g.Cells {
		for _, c := range r {
			s = s + strconv.Itoa(c)
		}
	}
	return s
}
