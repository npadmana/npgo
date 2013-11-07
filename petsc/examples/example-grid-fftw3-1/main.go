package main

import (
	"github.com/npadmana/petscgo"
	"github.com/npadmana/petscgo/grid/fftw3"
)

func main() {
	petscgo.Initialize()
	defer petscgo.Finalize()
	fftw3.Initialize()
	defer fftw3.Cleanup()

	rank, _ := petscgo.RankSize()

	dims := []int64{32, 16, 8, 12}

	// Test size routine
	lsize, n0, n1 := fftw3.LocalSizeTransposed(dims)
	petscgo.SyncPrintf("Rank %d : size = %d, n0=(%d,%d), n1=(%d,%d)\n", rank, lsize, n0[0], n0[1], n1[0], n1[1])
	petscgo.SyncFlush()

	// Ok, go create the grids
	dims = []int64{8, 8, 8}
	greal, gcmplx := fftw3.New(dims)
	petscgo.Printf("Considering the real array....\n")
	petscgo.SyncPrintf("Rank %d\n%s", rank, greal)
	petscgo.SyncFlush()

	petscgo.Printf("Considering the complex array....\n")
	petscgo.SyncPrintf("Rank %d\n%s", rank, gcmplx)
	petscgo.SyncFlush()

}
