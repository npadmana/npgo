package main

import (
	"github.com/npadmana/petscgo"
)

func main() {
	if err := petscgo.Initialize(); err != nil {
		petscgo.Fatal(err)
	}
	defer func() {
		if err := petscgo.Finalize(); err != nil {
			petscgo.Fatal(err)
		}
	}()
	rank, _ := petscgo.RankSize()

	// Create a vector using the local size
	v, err := petscgo.NewVec(5, petscgo.DETERMINE)
	if err != nil {
		petscgo.Fatal(err)
	}
	n1, err := v.LocalSize()
	if err != nil {
		petscgo.Fatal(err)
	}
	lo, hi, err := v.OwnRange()
	if err != nil {
		petscgo.Fatal(err)
	}
	petscgo.SyncPrintf("%d rank has local size %d [%d, %d]\n", rank, n1, lo, hi)
	petscgo.SyncFlush()
	err = v.Destroy()
	if err != nil {
		petscgo.Fatal(err)
	}

	// Create a vector using the local size
	v, err = petscgo.NewVec(petscgo.DECIDE, 100)
	if err != nil {
		petscgo.Fatal(err)
	}
	n1, err = v.LocalSize()
	if err != nil {
		petscgo.Fatal(err)
	}
	lo, hi, err = v.OwnRange()
	if err != nil {
		petscgo.Fatal(err)
	}
	petscgo.SyncPrintf("%d rank has local size %d [%d, %d]\n", rank, n1, lo, hi)
	petscgo.SyncFlush()
	err = v.Destroy()
	if err != nil {
		petscgo.Fatal(err)
	}

}
