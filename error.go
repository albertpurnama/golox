package main

import "fmt"

type CustomError struct {
	line    int
	where   string
	message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %s\nLine: %d\nWhere: %s", e.message, e.line, e.where)
}

func NewCustomError(line int, where string, message string) *CustomError {
	return &CustomError{line: line, where: where, message: message}
}
