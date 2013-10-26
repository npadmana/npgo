package cosmo

import (
	"math"
)

// Simple cosmology structure
type LCDM struct {
	h       float64 // little h
	om, ode float64 // physical densities in matter and dark energy
}

// NewFlatLCDMSimple returns a simple cosmology structure, taking in OmegaM0 and h
// and setting other values to default values.
func NewFlatLCDMSimple(OM0, h float64) (c LCDM) {
	c.h = h
	c.om = OM0 * h * h
	c.ode = (1 - OM0) * h * h
}

// Hubble method
func (c LCDM) Hubble
