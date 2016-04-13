package main

import (
	"errors"
	"regexp"
)

var (
	ErrIncorrectLength   = errors.New("Grid is incorrect length")
	ErrInvalidCharacters = errors.New("Grid contains invalid characters")
)

func isValidGrid(grid string) error {
	if len(grid) != 81 {
		return ErrIncorrectLength
	}
	return validChars(grid)
}

func validChars(grid string) error {
	validChars := regexp.MustCompile(`^(\d|\s)*$`)
	if !validChars.MatchString(grid) {
		return ErrInvalidCharacters
	}
	return nil
}
