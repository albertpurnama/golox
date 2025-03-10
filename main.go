package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	lox := Lox{}

	if len(args) > 1 {
		fmt.Println("Usage: go run main.go <command>")
		os.Exit(1)
	}

	if len(args) == 0 {
		lox.runPrompt()
		return
	}

	lox.runFile(args[0])
}

type Lox struct {
	hadError bool
}

func (l *Lox) runFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	loxErr := l.run(string(content))
	if loxErr != nil {
		l.report(loxErr)
	}

	if l.hadError {
		os.Exit(1)
	}
}

func (l *Lox) runPrompt() {
	fmt.Println("Initializing interpreter...")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		loxErr := l.run(input)
		if loxErr != nil {
			l.report(loxErr)
		}
	}
}

func (l *Lox) run(source string) *CustomError {
	scanner := Scanner{source: source}

	tokens, err := scanner.ScanTokens()
	if err != nil {
		return err
	}

	for idx, token := range tokens {
		fmt.Printf("%d: %s\n", idx, token)
	}

	return nil
}

func (l *Lox) report(err *CustomError) {
	fmt.Println("[line ", err.line, "] Error ", err.where, ": ", err.message)
	l.hadError = true
}
