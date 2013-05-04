package main

import (
	"fmt"
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

	// Create a vector using the global size
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

	// Try running ownershipranges
	if rank == 0 {
		rr, err := v.Ranges()
		if err != nil {
			petscgo.Fatal(err)
		}
		fmt.Println(rr)
	}

	// Set and then access the array
	if err := v.Set(3.1415926); err != nil {
		petscgo.Fatal(err)
	}
	if err := v.GetArray(); err != nil {
		petscgo.Fatal(err)
	}
	petscgo.SyncPrintf("%d rank has local size %d \n", rank, len(v.Arr))
	petscgo.SyncFlush()
	fmt.Println(v.Arr[0:2])
	if err := v.RestoreArray(); err != nil {
		petscgo.Fatal(err)
	}

	err = v.Destroy()
	if err != nil {
		petscgo.Fatal(err)
	}

}
