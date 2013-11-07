package main

import (
	"github.com/npadmana/npgo/petsc"
)

func main() {
	if err := petsc.Initialize(); err != nil {
		petsc.Fatal(err)
	}
	defer func() {
		if err := petsc.Finalize(); err != nil {
			petsc.Fatal(err)
		}
	}()
	rank, size := petsc.RankSize()

	petsc.Printf("Initialization successful\n")
	petsc.SyncPrintf("Hello from rank %d of %d\n", rank, size)
	petsc.SyncFlush()
}
