package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	// Run tests in parallel
	t.Parallel()

	// define a struct to hold all the tests
	type test_case struct {
		a, b float64
		want float64
	}

	test_cases := []test_case{
		{a: 2, b: 2, want: 4},
		{a: 0, b: 5, want: 5},
		{a: -1, b: 0, want: -1},
		{a: 1, b: -1, want: 0},
		{a: -1, b: -1, want: -2},
	}

	for _, test_case := range test_cases {
		got := calculator.Add(test_case.a, test_case.b)
		if got != test_case.want {
			t.Errorf("Addition of %.1f and %.1f resulted in %.1f, wanted %.1f", test_case.a, test_case.b, got, test_case.want)
		}
	}

}

func TestSubtract(t *testing.T) {
	// Run tests in parallel
	t.Parallel()

	// define a struct to hold all the tests
	type test_case struct {
		a, b float64
		want float64
	}

	test_cases := []test_case{
		{a: 0, b: 1, want: -1},
		{a: 10, b: 0, want: 10},
		{a: -1, b: -1, want: 0},
		{a: -2, b: -1, want: -1},
		{a: -2, b: -2, want: 0},
	}

	for _, test_case := range test_cases {
		got := calculator.Subtract(test_case.a, test_case.b)
		if got != test_case.want {
			t.Errorf("Subtraction of %.1f and %.1f resulted in %.1f, wanted %.1f", test_case.a, test_case.b, got, test_case.want)
		}
	}

}

func TestMultiply(t *testing.T) {
	// Run tests in parallel
	t.Parallel()

	// define a struct to hold all the tests
	type test_case struct {
		a, b float64
		want float64
	}

	test_cases := []test_case{
		{a: 0, b: 1, want: 0},
		{a: 10, b: 1, want: 10},
		{a: 6, b: 6, want: 36},
		{a: -2, b: -3, want: 6},
		{a: -1, b: 2, want: -2},
	}

	for _, test_case := range test_cases {
		got := calculator.Multiply(test_case.a, test_case.b)
		if got != test_case.want {
			t.Errorf("Multiplication of %.1f and %.1f resulted in %.1f, wanted %.1f", test_case.a, test_case.b, got, test_case.want)
		}
	}

}

func TestDivide(t *testing.T) {
	t.Parallel()

	// variables
	type test_case struct {
		a, b float64
		want float64
	}

	test_cases := []test_case{
		{a: 3, b: 3, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 1, b: 3, want: 0.333333},
	}

	for _, test_case := range test_cases {
		got, err := calculator.Divide(test_case.a, test_case.b)

		if err != nil {
			t.Fatalf("Divide (%f, %f): unexpected error status: %v", test_case.a, test_case.b, err)
		}

		if !CloseEnough(test_case.want, got, 0.001) {
			t.Errorf("Division of %f and %f resulted in %f, wanted %f", test_case.a, test_case.b, got, test_case.want)
		}
	}

}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("Should return divide by 0 error")
	}
}

func TestSquareRoot(t *testing.T) {
	t.Parallel()

	type test_case struct {
		a            float64
		want         float64
		err_expected bool
	}

	test_cases := []test_case{
		{a: 1, want: 1, err_expected: false},
		{a: -1, want: 999, err_expected: true},
		{a: 9, want: 3, err_expected: false},
	}

	for _, test_case := range test_cases {
		got, err := calculator.SquareRoot(test_case.a)
		err_recieved := (err != nil)

		if err_recieved != test_case.err_expected {
			t.Fatalf("SquareRoot (%.1f): unexpected error status: %v", test_case.a, err)
		}

		if !test_case.err_expected && test_case.want != got {
			t.Errorf("Square Root (%.1f): want %.1f, got %.1f", test_case.a, test_case.want, got)
		}

	}

}

func CloseEnough(a, b, tolerance float64) bool {
	// this is used to set the tolerance of two numbers being considered equal
	return math.Abs(a-b) <= tolerance
}
