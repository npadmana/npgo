package main

import (
	"github.com/npadmana/npgo/petsc"
	"github.com/npadmana/npgo/petsc/grid/fftw3"
)

func main() {
	petsc.Initialize()
	defer petsc.Finalize()
	fftw3.Initialize()
	defer fftw3.Cleanup()

	rank, _ := petsc.RankSize()

	dims := []int64{32, 16, 8, 12}

	// Test size routine
	lsize, n0, n1 := fftw3.LocalSizeTransposed(dims)
	petsc.SyncPrintf("Rank %d : size = %d, n0=(%d,%d), n1=(%d,%d)\n", rank, lsize, n0[0], n0[1], n1[0], n1[1])
	petsc.SyncFlush()

	// Ok, go create the grids
	dims = []int64{8, 8, 8}
	greal, gcmplx := fftw3.New(dims)
	petsc.Printf("Considering the real array....\n")
	petsc.SyncPrintf("Rank %d\n%s", rank, greal)
	petsc.SyncFlush()

	petsc.Printf("Considering the complex array....\n")
	petsc.SyncPrintf("Rank %d\n%s", rank, gcmplx)
	petsc.SyncFlush()

}
