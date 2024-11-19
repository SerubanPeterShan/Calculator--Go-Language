package calculator

import (
	"errors"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func square(a int) int {
	return a * a
}
