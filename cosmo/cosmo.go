// Package cosmo contains basic cosmology routines, and is the
// top-level package for others.
package cosmo

import (
	"math"

	"github.com/npadmana/npgo/gsl"
	"github.com/npadmana/npgo/gsl/integ"
)

const (
	CLight = 299792.458 // in km/s
)

func A2Z(a float64) float64 {
	return (1 / a) - 1
}

func Z2A(z float64) float64 {
	return 1 / (1 + z)
}

// Interface Hubbler defines a function Hubble that returns h(a), defined such that
// the Hubble parameter is 100 Hubble(a) km/s/Mpc
type Hubbler interface {
	Hubble(a float64) float64
}

// Func SinK computes the curvature dependent distance
func SinK(omkh2, d float64) float64 {
	k := math.Sqrt(math.Abs(omkh2))
	kd := k * d
	var ret float64
	switch {
	case (omkh2 > 0) && (kd > 1.e-2):
		ret = math.Sinh(kd) / k
	case (omkh2 < 0) && (kd > 1.e-2):
		ret = math.Sin(kd) / k
	case (omkh2 >= 0) && (kd < 1.e-2):
		ret = d + kd*kd*d/6
	case (omkh2 < 0) && (kd < 1.e-2):
		ret = d - kd*kd*d/6
	}
	return ret
}

// Func ComDis(Hubbler, avals) computes the comoving distance in Mpc
//
// This function panics if the integrator failed for some reason, but throws away the
// estimated error.
func ComDis(h Hubbler, avals []float64) []float64 {
	retval := make([]float64, len(avals))
	ff := func(a float64) float64 { return 1 / (a * a * h.Hubble(a)) }
	w := integ.NewWork(1000)
	defer w.Free()
	var res gsl.Result
	var err error
	for i, a := range avals {
		res, err = integ.Qags(ff, gsl.Interval{a, 1}, gsl.Eps{1e-7, 1e-7}, w)
		if err != nil {
			panic(err)
		}
		retval[i] = res.Res * CLight / 100
	}
	return retval
}
