package calculator

/*
Testing in Go: NO external framework needed - the standard library's
`testing` package + `go test` command is enough for most projects.
Convention: code in foo.go is tested by foo_test.go in the SAME package.
Run with:  go test ./...        (add -v for verbose, -cover for coverage)
*/

import "errors"

func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
