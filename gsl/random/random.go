// Package random wraps the GSL random number generators
package random

/* 
#cgo pkg-config: gsl


#include "gsl/gsl_rng.h"
*/
import "C"

import (
	"errors"
)

type RNGType int

const (
	MT19937 RNGType = iota
)

// RNG wraps the GSL random number generators
type RNG struct {
	rng *C.gsl_rng
}

func convertRNGType(r RNGType) (*C.gsl_rng_type, error) {
	var ret *C.gsl_rng_type
	var err error
	switch r {
	case MT19937:
		ret, err = C.gsl_rng_mt19937, nil
	default:
		ret, err = C.gsl_rng_mt19937, errors.New("Unknown random number generator")
	}
	return ret, err
}

// New returns a new RNG of type r
func New(r RNGType) (*RNG, error) {
	ret := new(RNG)
	rtype, err := convertRNGType(r)
	if err != nil {
		return nil, err
	}
	ret.rng = C.gsl_rng_alloc(rtype)
	return ret, nil
}

// NewMT returns a Mersenne-Twister RNG
func NewMT() (*RNG, error) {
	return New(MT19937)
}

// Free cleans up the RNG
func (r *RNG) Free() {
	C.gsl_rng_free(r.rng)
}

// Seed seeds the random number generator
func (r *RNG) Seed(s int64) {
	C.gsl_rng_set(r.rng, C.ulong(s))
}

// Get returns a random integer
func (r *RNG) Get() int64 {
	return int64(C.gsl_rng_get(r.rng))
}

// Uniform returns a uniform random number between [0,1)
func (r *RNG) Uniform() float64 {
	return float64(C.gsl_rng_uniform(r.rng))
}

// UniformPos returns a uniform random number between (0,1)
func (r *RNG) UniformPos() float64 {
	return float64(C.gsl_rng_uniform_pos(r.rng))
}

// UniformInt returns a uniform random number between (0,1)
func (r *RNG) UniformInt(max int64) int64 {
	return int64(C.gsl_rng_uniform_int(r.rng, C.ulong(max)))
}
