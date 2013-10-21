// Package deriv wraps the numerical differentiation routines.
package deriv

/*
#cgo pkg-config: gsl

#include <gsl/gsl_deriv.h>

extern double derivCB(double x, void *params);

static gsl_function mkderivCB(void *data) {
	gsl_function gf;
	gf.function = derivCB;
	gf.params = data;
	return gf;
}

*/
import "C"

import (
	"errors"
	"github.com/npadmana/npgo/gsl"
	"unsafe"
)

//export derivCB
func derivCB(x C.double, data unsafe.Pointer) C.double {
	ff := (*gsl.GSLFuncWrapper)(data)
	return C.double(ff.Gofunc(float64(x)))
}

// Different types of derivatives
type DerivType int

const (
	Backward DerivType = iota - 1
	Central
	Forward
)

// Diff computes the derivative of ff, returns derivative and an error
func Diff(dir DerivType, ff gsl.F, x, h float64) (gsl.Result, error) {
	var y, err C.double
	var ret C.int
	var gf C.gsl_function

	data := gsl.GSLFuncWrapper{ff}
	gf = C.mkderivCB(unsafe.Pointer(&data))
	switch dir {
	case Central:
		ret = C.gsl_deriv_central(&gf, C.double(x), C.double(h), &y, &err)
	case Forward:
		ret = C.gsl_deriv_forward(&gf, C.double(x), C.double(h), &y, &err)
	case Backward:
		ret = C.gsl_deriv_backward(&gf, C.double(x), C.double(h), &y, &err)
	default:
		panic(errors.New("Unknown direction"))
	}
	if ret != 0 {
		return gsl.Result{float64(y), float64(err)}, gsl.Errno(ret)
	}
	return gsl.Result{float64(y), float64(err)}, nil
}
