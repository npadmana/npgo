package main

import (
	"fmt"
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

	// Create a vector using the local size
	v, err := petsc.NewVec(5, petsc.DETERMINE)
	if err != nil {
		petsc.Fatal(err)
	}
	n1, err := v.LocalSize()
	if err != nil {
		petsc.Fatal(err)
	}
	lo, hi, err := v.OwnRange()
	if err != nil {
		petsc.Fatal(err)
	}
	petsc.SyncPrintf("%d rank has local size %d [%d, %d]\n", rank, n1, lo, hi)
	petsc.SyncFlush()
	err = v.Destroy()
	if err != nil {
		petsc.Fatal(err)
	}

	// Create a vector using the global size
	v, err = petsc.NewVec(petsc.DECIDE, 100)
	if err != nil {
		petsc.Fatal(err)
	}
	n1, err = v.LocalSize()
	if err != nil {
		petsc.Fatal(err)
	}
	lo, hi, err = v.OwnRange()
	if err != nil {
		petsc.Fatal(err)
	}
	petsc.SyncPrintf("%d rank has local size %d [%d, %d]\n", rank, n1, lo, hi)
	petsc.SyncFlush()

	// Set and then access the array
	if err := v.Set(3.1415926); err != nil {
		petsc.Fatal(err)
	}

	// Try running ownershipranges
	if rank == 0 {
		rr, err := v.Ranges()
		if err != nil {
			petsc.Fatal(err)
		}
		fmt.Println(rr)
		if size > 2 {
			ix := []int64{rr[1], rr[2], rr[3]}
			y := []float64{4.14, 5.14, 6.14}
			v.SetValues(ix, y, true)
		}
	}
	v.AssemblyBegin()
	v.AssemblyEnd()

	if err := v.GetArray(); err != nil {
		petsc.Fatal(err)
	}
	petsc.SyncPrintf("%d rank has local size %d \n", rank, len(v.Arr))
	petsc.SyncFlush()
	fmt.Println(rank, v.Arr[0:2])
	if err := v.RestoreArray(); err != nil {
		petsc.Fatal(err)
	}

	sum, _ := v.Sum()
	petsc.Printf("Sum = %f\n", sum)
	max, _, _ := v.Max()
	petsc.Printf("Max = %f\n", max)
	min, _, _ := v.Min()
	petsc.Printf("Max = %f\n", min)
	v.Scale(0.3)
	sum, _ = v.Sum()
	petsc.Printf("Sum = %f\n", sum)

	err = v.Destroy()
	if err != nil {
		petsc.Fatal(err)
	}

}
