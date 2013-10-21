// Package spline provides access to the GSL spline functions.
package spline

/*
#cgo pkg-config: gsl


#include "gsl/gsl_spline.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"github.com/npadmana/npgo/gsl"
)

// SplineType defines the various types of splines available
type SplineType int

const (
	Linear SplineType = iota
	Polynomial
	Cubic
	CubicPeriodic
	Akima
	AkimaPeriodic
)

func convertSplineType(s SplineType) (*C.gsl_interp_type, error) {
	var ret *C.gsl_interp_type
	var err error
	switch s {
	case Linear:
		ret, err = C.gsl_interp_linear, nil
	case Polynomial:
		ret, err = C.gsl_interp_polynomial, nil
	case Cubic:
		ret, err = C.gsl_interp_cspline, nil
	case CubicPeriodic:
		ret, err = C.gsl_interp_cspline_periodic, nil
	case Akima:
		ret, err = C.gsl_interp_akima, nil
	case AkimaPeriodic:
		ret, err = C.gsl_interp_akima_periodic, nil
	default:
		ret, err = C.gsl_interp_linear, errors.New("Unknown spline type")
	}
	return ret, err
}

// Spline wraps the GSL spline and accelerator structures
type Spline struct {
	sp  *C.gsl_spline
	acc *C.gsl_interp_accel
}

// Free frees the spline variables
func (s *Spline) Free() {
	C.gsl_interp_accel_free(s.acc)
	C.gsl_spline_free(s.sp)
}

// NewSpline creates a new Spline struct
func New(s SplineType, xa, ya []float64) (*Spline, error) {

	// Check inputs
	nx := len(xa)
	ny := len(ya)
	if nx != ny {
		return nil, fmt.Errorf("Incompatible dimensions in NewSpline: x(%d) != y(%d)", nx, ny)
	}

	// Convert type of spline
	sptype, err := convertSplineType(s)
	if err != nil {
		return nil, err
	}

	// Create a new object
	sp := new(Spline)
	sp.sp = C.gsl_spline_alloc(sptype, C.size_t(nx))
	sp.acc = C.gsl_interp_accel_alloc()

	// Initialize the spline object
	ret := C.gsl_spline_init(sp.sp, (*C.double)(&xa[0]), (*C.double)(&ya[0]), C.size_t(nx))
	if ret != 0 {
		return nil, gsl.Errno(ret)
	}

	return sp, nil
}

// Eval evaluates the spline at x.
// If x is out of bounds, the code will panic.
func (s *Spline) Eval(x float64) (float64, error) {
	var y C.double
	ret := C.gsl_spline_eval_e(s.sp, C.double(x), s.acc, &y)
	if ret != 0 {
		return float64(y), gsl.Errno(ret)
	}
	return float64(y), nil
}

// Deriv evaluates the derivative of the spline at x.
// If x is out of bounds, the code will panic.
func (s *Spline) Deriv(x float64) (float64, error) {
	var y C.double
	ret := C.gsl_spline_eval_deriv_e(s.sp, C.double(x), s.acc, &y)
	if ret != 0 {
		return float64(y), gsl.Errno(ret)
	}
	return float64(y), nil
}

// Integrate evaluates the integral of the spline from lo to hi.
// If this put it out of bounds, x
// If x is out of bounds, the code will panic.
func (s *Spline) Integrate(lo, hi float64) (float64, error) {
	var y C.double
	ret := C.gsl_spline_eval_integ_e(s.sp, C.double(lo), C.double(hi), s.acc, &y)
	if ret != 0 {
		return float64(y), gsl.Errno(ret)
	}
	return float64(y), nil
}
