package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) (float64, error) {
	// Could there be a case that an addition error could occur?
	return a + b, nil
}

func Subtract(a, b float64) (float64, error) {
	// Could there be a case that an subtraction error could occur?
	return a - b, nil
}

func Multiply(a, b float64) (float64, error) {
	// Could there be a case that an multiplication error could occur?
	return a * b, nil
}

func Divide(a, b float64) (float64, error) {
	// Could there be a case that an division error could occur? - Yes, if you try and divide by zero
	if b == 0 {
		return 0, errors.New("Divide by Zero not allowed")
	}
	return a / b, nil
}

func SquareRoot(a float64) (float64, error) {
	// Could there be a case where multiplying two numbers would be an invalid square root? - Yes, if you try and use a negative number
	// Signbit() returns true if the number is a negative
	if math.Signbit(a) {
		return 0, errors.New("Square Root must use real numbers")
	}
	return a * a, nil
}
