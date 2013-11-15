package sf

/*
#cgo pkg-config: gsl


#include "gsl/gsl_sf_bessel.h"
*/
import "C"

import (
	"github.com/npadmana/npgo/gsl"
)

// BesselJ returns Jn(x)
func BesselJ(n int, x float64) float64 {
	var y C.double
	switch n {
	case 0:
		y = C.gsl_sf_bessel_J0(C.double(x))
	case 1:
		y = C.gsl_sf_bessel_J1(C.double(x))
	default:
		y = C.gsl_sf_bessel_Jn(C.int(n), C.double(x))
	}
	return float64(y)
}

// BesselJArr returns an array of Jn(x) where n runs from nmin to nmax inclusive
func BesselJArr(nmin, nmax int, x float64) []float64 {
	arr := make([]float64, nmax-nmin+1)
	ret := C.gsl_sf_bessel_Jn_array(C.int(nmin), C.int(nmax), C.double(x), (*C.double)(&arr[0]))
	if ret != 0 {
		panic(gsl.Errno(ret))
	}
	return arr
}

// SphBessel returns the spherical bessel function jl(x)
func SphBessel(l int, x float64) float64 {
	var y C.double
	switch l {
	case 0:
		y = C.gsl_sf_bessel_j0(C.double(x))
	case 1:
		y = C.gsl_sf_bessel_j1(C.double(x))
	case 2:
		y = C.gsl_sf_bessel_j2(C.double(x))
	default:
		y = C.gsl_sf_bessel_jl(C.int(l), C.double(x))
	}
	return float64(y)
}

// SphBesselArr returns an array of Jl(x) where l runs from 0 to nmax inclusive
//
// Note that GSL has two implementations; we use the default one, not the one based
// on Steed's algorithm.
func SphBesselArr(lmax int, x float64) []float64 {
	arr := make([]float64, lmax+1)
	ret := C.gsl_sf_bessel_jl_array(C.int(lmax), C.double(x), (*C.double)(&arr[0]))
	if ret != 0 {
		panic(gsl.Errno(ret))
	}
	return arr
}
