package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {

	return a + b

}

func Subtract(a, b float64) float64 {

	return a - b

}

func Multiply(a, b float64) float64 {

	return a * b
}

func Divide(a, b float64) (float64, error) {
	// Catch divide by zero and return an error
	if b == 0 {
		return 0, errors.New("Divide by Zero not allowed")
	}
	return a / b, nil

}

func SquareRoot(a float64) (float64, error) {
	// Catch negative input
	// Signbit() returns true if the number is a negative
	if math.Signbit(a) {
		return 0, errors.New("Square Root must use real numbers")
	}
	return math.Sqrt(a), nil

}
