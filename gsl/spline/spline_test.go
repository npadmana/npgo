package spline

import (
	"fmt"
	"math"
	"testing"
)

func TestSplineLinear(t *testing.T) {
	xa := []float64{1, 2, 3, 4, 5}
	ya := []float64{1, 2, 3, 4, 5}

	sp, err := New(Linear, xa, ya)
	if err != nil {
		fmt.Println(err)
		t.Error("Error occurred allocating a spline")
	}
	defer sp.Free()

	var x, y float64
	x = 2.3
	y, err = sp.Eval(x)
	if err != nil {
		t.Errorf("Unexpected error returned : %v", err)
	}
	if math.Abs(x-y) > 1.e-7 {
		t.Errorf("Incorrect value: expected = %f, actual = %f", x, y)
	}

	// Test out of domain
	x = 5.7
	y, err = sp.Eval(x)
	if err == nil {
		t.Error("Expected an error, none reported")
	}

	// Test derivative
	x = 2.3
	y, err = sp.Deriv(x)
	if err != nil {
		t.Errorf("Unexpected error in derivative : %v", err)
	}
	if math.Abs(y-1) > 1.e-7 {
		t.Errorf("Incorrect value: expected = %f, actual = %f", 1, y)
	}

	// Test integrate
	y, err = sp.Integrate(2, 4)
	if err != nil {
		t.Errorf("Unexpected error in derivative : %v", err)
	}
	if math.Abs(y-6) > 1.e-7 {
		t.Errorf("Incorrect value: expected = %f, actual = %f", 6, y)
	}

}

func TestSplineCubic1(t *testing.T) {
	xa := []float64{1, 2, 3, 4, 5}
	ya := []float64{1, 2, 3, 4, 5}

	sp, err := New(Cubic, xa, ya)
	if err != nil {
		fmt.Println(err)
		t.Error("Error occurred allocating a spline")
	}
	defer sp.Free()

	var x, y float64
	x = 2.3
	y, err = sp.Eval(x)
	if err != nil {
		t.Errorf("Unexpected error returned : %v", err)
	}
	if math.Abs(x-y) > 1.e-7 {
		t.Errorf("Incorrect value: expected = %f, actual = %f", x, y)
	}

}

func TestSplineCubic2(t *testing.T) {
	xa := []float64{1, 2, 3, 4, 5}
	ya := []float64{1, 4, 9, 16, 25}

	sp, err := New(Cubic, xa, ya)
	if err != nil {
		fmt.Println(err)
		t.Error("Error occurred allocating a spline")
	}
	defer sp.Free()

	var x, y float64
	x = 3.1
	y0 := x * x
	y, err = sp.Eval(x)
	if err != nil {
		t.Errorf("Unexpected error returned : %v", err)
	}
	if math.Abs(y0-y) > 1.e-2 {
		t.Errorf("Incorrect value: expected = %f, actual = %f", y0, y)
	}

}
