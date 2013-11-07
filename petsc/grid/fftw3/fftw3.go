package fftw3

/*
#cgo pkg-config: fftw3-mpi ompi

#include <stddef.h>
#include "fftw3-mpi.h"

int size_ptrdiff() {
	return (int) sizeof(ptrdiff_t);
}

*/
import "C"

import (
	"errors"
	"unsafe"

	"github.com/npadmana/npgo/petsc"
)

// Initialize initializes FFTW3 for MPI usage
func Initialize() {
	if C.size_ptrdiff() != 8 {
		petsc.Fatal(errors.New("size(ptrdiff_t) != 8"))
	}

	// Initialize FFTW3
	C.fftw_mpi_init()
}

// Cleanup cleans up FFTW3
func Cleanup() {
	C.fftw_mpi_cleanup()
}

// LocalSizeTransposed returns the local size required for the transform,
// as well as the start and end positions for the transposed and untransposed
// dimensions.
//
// Unlike the FFTW3 interface, dims here are the dimensions for the real transform.
//
// returns the local size (number of real elts), n0 and n1 start and local sizes.
func LocalSizeTransposed(dims []int64) (int64, [2]int64, [2]int64) {
	ndim := len(dims)
	if ndim < 2 {
		petsc.Fatal(errors.New("Need at least two dimensions"))
	}

	rnk := C.int(ndim)
	// We're going to change the dimensions, so make a copy.
	dims1 := make([]int64, ndim)
	copy(dims1, dims)
	dims1[ndim-1] = dims[ndim-1]/2 + 1

	// ptrdiff_t fftw_mpi_local_size_transposed(int rnk, const ptrdiff_t *n, MPI_Comm comm,
	//                                               ptrdiff_t *local_n0, ptrdiff_t *local_0_start,
	//                                               ptrdiff_t *local_n1, ptrdiff_t *local_1_start);

	var n0, n0start, n1, n1start C.ptrdiff_t
	lsize := C.fftw_mpi_local_size_transposed(rnk, (*C.ptrdiff_t)(unsafe.Pointer(&dims1[0])), petsc.WORLD,
		&n0, &n0start, &n1, &n1start)

	ln0 := [2]int64{int64(n0), int64(n0start)}
	ln1 := [2]int64{int64(n1), int64(n1start)}

	return int64(lsize) * 2, ln0, ln1
}
