package main

import (
	"errors"
	"fmt"
	"os"
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

func main() {
	gridInput := os.Args[1]

	if err := valid(gridInput); err != nil {
		fmt.Println(err)
		return
	}

	grid, err := transformIntoGrid(gridInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Solving...")
	fmt.Println(grid.toString())

	if err = grid.simpleSolve(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Solution...")
	fmt.Println(grid.toString())
}

func valid(gridInput string) error {
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

type grid struct {
	Cells [][]int
}

func transformIntoGrid(gridInput string) (*grid, error) {
	g := grid{
		Cells: [][]int{},
	}

	if err := g.addRows(gridInput); err != nil {
		return nil, err
	}

	if g.hasDuplicates() {
		return nil, ErrDuplicateValues
	}

	return &g, nil
}

func (g *grid) addRows(gridInput string) error {
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

type selectionType int

const (
	colSelection selectionType = iota
	rowSelection
	subMatrixSelection
)

func (g *grid) getSelection(stype selectionType, Location int) []int {
	selection := []int{}

	switch stype {
	case colSelection:
		for _, s := range g.Cells {
			selection = append(selection, s[Location])
		}
	case rowSelection:
		selection = g.Cells[Location]
	case subMatrixSelection:
		selection = g.subMatrices()[Location]
	}

	return selection
}

func (g *grid) hasDuplicates() bool {
	selectionTypes := []selectionType{colSelection, rowSelection, subMatrixSelection}

	for i, _ := range g.Cells {
		for _, stype := range selectionTypes {
			selection := g.getSelection(stype, i)
			if hasDuplicateCells(selection) {
				return true
			}
		}
	}
	return false
}

func hasDuplicateCells(cells []int) bool {
	for i, cell := range cells {
		if cell != 0 {
			for j := i + 1; j < len(cells); j++ {
				if cell == cells[j] {
					return true
				}
			}
		}
	}
	return false
}

func (g *grid) subMatrices() [][]int {
	submatrices := make([][]int, 9)
	for i := 0; i < len(submatrices); i += 3 {
		for j, _ := range g.Cells {
			if j >= i && j < i+3 {
				submatrices[i] = append(submatrices[i], g.Cells[j][0:3]...)
				submatrices[i+1] = append(submatrices[i+1], g.Cells[j][3:6]...)
				submatrices[i+2] = append(submatrices[i+2], g.Cells[j][6:9]...)
			}
		}
	}
	return submatrices
}

type unsolvedCell struct {
	row    int
	column int
}

func (g *grid) simpleSolve() error {
	unsolved := g.unsolved()
	err := g.findValidNumbers(unsolved, 0)
	if err != nil {
		return err
	}
	return nil
}

func (g *grid) unsolved() []unsolvedCell {
	unsolved := []unsolvedCell{}
	for i, row := range g.Cells {
		for j, s := range row {
			if s == 0 {
				u := unsolvedCell{
					row:    i,
					column: j,
				}
				unsolved = append(unsolved, u)
			}
		}
	}
	return unsolved
}

func (g *grid) findValidNumbers(unsolved []unsolvedCell, current int) error {
	if g.isSolved() {
		return nil
	}

	cell := unsolved[current]
	if err := g.increase(cell.row, cell.column); err != nil {
		if current == 0 {
			return ErrNoSolution
		}

		g.reset(cell.row, cell.column)
		return g.findValidNumbers(unsolved, current-1)
	}

	if current != (len(unsolved) - 1) {
		return g.findValidNumbers(unsolved, current+1)
	}

	return nil
}

func (g *grid) increase(r, c int) error {
	if g.Cells[r][c] == 9 {
		return ErrReachedLimit
	}

	g.Cells[r][c]++
	if g.hasDuplicates() {
		return g.increase(r, c)
	}

	return nil
}

func (g *grid) reset(r, c int) {
	g.Cells[r][c] = 0
}

func (g *grid) isSolved() bool {
	if len(g.unsolved()) == 0 {
		if !g.hasDuplicates() {
			return true
		}
	}

	return false
}

func (g *grid) toString() string {
	var s string
	for _, r := range g.Cells {
		for _, c := range r {
			s = fmt.Sprintf("%s %s", s, strconv.Itoa(c))
		}
		s = s + "\n"
	}
	return s
}
