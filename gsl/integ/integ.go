// Package integ wraps the basic integration routines
package integ

/*
#cgo pkg-config: gsl

#include <gsl/gsl_integration.h>

extern double integCB(double x, void *params);

static gsl_function mkintegCB(void *data) {
	gsl_function gf;
	gf.function = integCB;
	gf.params = data;
	return gf;
}

*/
import "C"

import (
	"github.com/npadmana/npgo/gsl"
	"math"
	"unsafe"
)

//export integCB
func integCB(x C.double, data unsafe.Pointer) C.double {
	ff := (*gsl.GSLFuncWrapper)(data)
	return C.double(ff.Gofunc(float64(x)))
}

// Workspace is the GSL integration workspace
type WorkSpace struct {
	n int
	w *C.gsl_integration_workspace
}

// NewWork allocate new workspace
func NewWork(n int) *WorkSpace {
	ret := new(WorkSpace)
	ret.n = n
	ret.w = C.gsl_integration_workspace_alloc(C.size_t(n))
	return ret
}

// Free frees workspace w
func (w *WorkSpace) Free() {
	C.gsl_integration_workspace_free(w.w)
}

func Qags(ff gsl.F, ab gsl.Interval, eps gsl.Eps, w *WorkSpace) (gsl.Result, error) {
	// Make a gsl_function
	var gf C.gsl_function
	data := gsl.GSLFuncWrapper{ff}
	gf = C.mkintegCB(unsafe.Pointer(&data))

	// Check to see if we have a positive/-negative infinity
	pinf := math.IsInf(ab.Hi, 1)
	ninf := math.IsInf(ab.Lo, -1)

	var ret C.int
	var y, err C.double
	// Switch on options
	switch {
	case pinf && ninf:
		ret = C.gsl_integration_qagi(&gf, C.double(eps.Abs), C.double(eps.Rel), C.size_t(w.n), w.w, &y, &err)
	case pinf:
		ret = C.gsl_integration_qagiu(&gf, C.double(ab.Lo), C.double(eps.Abs), C.double(eps.Rel), C.size_t(w.n), w.w, &y, &err)
	case ninf:
		ret = C.gsl_integration_qagil(&gf, C.double(ab.Hi), C.double(eps.Abs), C.double(eps.Rel), C.size_t(w.n), w.w, &y, &err)
	default:
		ret = C.gsl_integration_qags(&gf, C.double(ab.Lo), C.double(ab.Hi), C.double(eps.Abs), C.double(eps.Rel), C.size_t(w.n), w.w, &y, &err)
	}
	if ret != 0 {
		return gsl.Result{float64(y), float64(err)}, gsl.Errno(ret)
	}
	return gsl.Result{float64(y), float64(err)}, nil

}
