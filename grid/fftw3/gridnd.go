// Package fftw3 is an implementation of an MPI/FFTW/PETSc 3D grid.
//
package fftw3

import (
	"errors"
	"fmt"

	"github.com/npadmana/petscgo"
)

type Grid struct {
	V             *petscgo.Vec // We store the array here
	dims, strides []int64      // Dimension information
	lo, hi        []int64      // The lo-hi dimensions before transposition
}

// String returns grid information
func (g *Grid) String() string {
	s := fmt.Sprintf("Dimensions : %v\nStride : %v\nLo : %v\nHi : %v\n", g.dims, g.strides, g.lo, g.hi)
	return s
}

// Ndim returns the number of dimensions
func (g *Grid) Ndim() int {
	return len(g.dims)
}

// Dimensions returns the full dimensions.
func (g *Grid) Dimensions() []int64 {
	return g.dims
}

// Strides returns the strides
func (g *Grid) Strides() []int64 {
	return g.strides
}

// Lo returns the lower indices
func (g *Grid) Lo() []int64 {
	return g.lo
}

// Hi returns the upper indices, exclusive
func (g *Grid) Hi() []int64 {
	return g.hi
}

// helper function
// complex is assumed to be transposed.
func simpleNew(dims []int64, n0 [2]int64, complex bool) *Grid {
	ndim := len(dims)
	g := new(Grid)
	g.dims = make([]int64, ndim)
	copy(g.dims, dims)

	// This is the size in the complex dimension
	g.dims[ndim-1] = g.dims[ndim-1]/2 + 1

	// Now fiddle things if complex
	if complex {
		// Transpose leading dimensions
		g.dims[0], g.dims[1] = g.dims[1], g.dims[0]
	} else {
		// Padding
		g.dims[ndim-1] *= 2
	}

	// set strides
	g.strides = make([]int64, ndim)
	g.strides[ndim-1] = 1
	for i := ndim - 2; i >= 0; i-- {
		g.strides[i] = g.strides[i+1] * g.dims[i+1]
	}

	// Set lo and hi terms
	g.lo = make([]int64, ndim)
	g.hi = make([]int64, ndim)
	for i := range g.lo {
		g.lo[i] = 0
		g.hi[i] = dims[i]
	}
	// Fix the regions that are split.
	if complex {
		g.hi[ndim-1] = g.dims[ndim-1]
		g.lo[1] = n0[1]
		g.hi[1] = g.lo[1] + n0[0]
	} else {
		g.lo[0] = n0[1]
		g.hi[0] = g.lo[0] + n0[0]
	}

	// Transpose if complex, fix trailing dimension if real
	if complex {
		g.dims[0], g.dims[1] = g.dims[1], g.dims[0]
	} else {
		g.dims[ndim-1] = dims[ndim-1]
	}

	return g
}

// New returns a new fftw3.Grid of sides dim[] (specified in configuration space)
func New(dims []int64) (greal *Grid, gcmplx *Grid) {
	// Check that dims is big enough
	ndim := len(dims)
	if ndim < 2 {
		petscgo.Fatal(errors.New("The grid must at least be 2D"))
	}

	g := new(Grid)
	g.dims = make([]int64, ndim)
	copy(g.dims, dims)

	// Figure out local dimensions
	lsize, n0, n1 := LocalSizeTransposed(dims)

	// Allocate the real and complex grids
	greal = simpleNew(dims, n0, false)
	gcmplx = simpleNew(dims, n1, true)

	// Ok, now allocate arrays and store the same in both places

	_ = lsize
	return
}
