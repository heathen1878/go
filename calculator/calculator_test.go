package calculator_test

import (
	"calculator"
	"testing"
)

func TestOperator(t *testing.T) {
	t.Parallel()

	// variables
	type test_case struct {
		fn           func(float64, float64) (float64, error)
		a, b         float64
		want         float64
		err_expected bool
	}

	test_cases := []test_case{
		{fn: calculator.Add, a: 2, b: 2, want: 4, err_expected: false},
		{fn: calculator.Add, a: 0, b: 5, want: 5, err_expected: false},
		{fn: calculator.Add, a: -1, b: 0, want: -1, err_expected: false},
		{fn: calculator.Add, a: 1, b: -1, want: 0, err_expected: false},
		{fn: calculator.Add, a: -1, b: -1, want: -2, err_expected: false},
		{fn: calculator.Subtract, a: 0, b: 1, want: -1, err_expected: false},
		{fn: calculator.Subtract, a: 10, b: 0, want: 10, err_expected: false},
		{fn: calculator.Subtract, a: -1, b: -1, want: 0, err_expected: false},
		{fn: calculator.Subtract, a: -2, b: -1, want: -1, err_expected: false},
		{fn: calculator.Subtract, a: -2, b: -2, want: 0, err_expected: false},
		{fn: calculator.Multiply, a: 0, b: 1, want: 0, err_expected: false},
		{fn: calculator.Multiply, a: 10, b: 1, want: 10, err_expected: false},
		{fn: calculator.Multiply, a: 6, b: 6, want: 36, err_expected: false},
		{fn: calculator.Multiply, a: -2, b: -3, want: 6, err_expected: false},
		{fn: calculator.Multiply, a: -1, b: 2, want: -2, err_expected: false},
		{fn: calculator.Divide, a: 3, b: 0, want: 999, err_expected: true},
		{fn: calculator.Divide, a: 3, b: 3, want: 1, err_expected: false},
	}

	for _, test_case := range test_cases {
		got, err := test_case.fn(test_case.a, test_case.b)
		err_recieved := (err != nil)

		if err_recieved != test_case.err_expected {
			t.Fatalf("Divide (%.1f, %.1f): unexpected error status: %v", test_case.a, test_case.b, err)
		}

		if !test_case.err_expected && test_case.want != got {
			t.Errorf("(%.1f, %.1f): Want %.1f, got %.1f", test_case.a, test_case.b, test_case.want, got)
		}
	}

}

func TestSquareRoot(t *testing.T) {
	t.Parallel()

	type test_case struct {
		fn           func(float64) (float64, error)
		a            float64
		want         float64
		err_expected bool
	}

	test_cases := []test_case{
		{fn: calculator.SquareRoot, a: 1, want: 1, err_expected: false},
		{fn: calculator.SquareRoot, a: -1, want: 999, err_expected: true},
		{fn: calculator.SquareRoot, a: 3, want: 9, err_expected: false},
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
