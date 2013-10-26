// Package cosmo contains basic cosmology routines, and is the
// top-level package for others.
package cosmo

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
	Hubble(a float64)
}
