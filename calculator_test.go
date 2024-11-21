package calculator

import (
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if result != 5 || err != nil {
		t.Errorf("Divide(10, 2) failed. Got %v, %v", result, err)
	}

	result, err = Divide(10, 0)
	if result != 0 || err == nil {
		t.Errorf("Divide(10, 0) failed. Got result: %v, error: %v", result, err)
	}
}

func TestSquare(t *testing.T) {
	result := square(10)
	if result != 100 {
		t.Errorf("Square(10) failed. Got %v, expected 100", result)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Divide(10, 2)
	}
}

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = square(10)
	}
}
