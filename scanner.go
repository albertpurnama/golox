package main

type Scanner struct {
	source string
}

type Token struct{}

func (s *Scanner) ScanTokens() ([]Token, *CustomError) {
	return []Token{
		{},
		{},
		{},
		{},
	}, nil
}
