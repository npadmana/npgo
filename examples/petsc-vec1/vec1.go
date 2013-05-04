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
	rank, size := petscgo.RankSize()

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

	// Set and then access the array
	if err := v.Set(3.1415926); err != nil {
		petscgo.Fatal(err)
	}

	// Try running ownershipranges
	if rank == 0 {
		rr, err := v.Ranges()
		if err != nil {
			petscgo.Fatal(err)
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
		petscgo.Fatal(err)
	}
	petscgo.SyncPrintf("%d rank has local size %d \n", rank, len(v.Arr))
	petscgo.SyncFlush()
	fmt.Println(rank, v.Arr[0:2])
	if err := v.RestoreArray(); err != nil {
		petscgo.Fatal(err)
	}

	sum, _ := v.Sum()
	petscgo.Printf("Sum = %f\n", sum)
	max, _, _ := v.Max()
	petscgo.Printf("Max = %f\n", max)
	min, _, _ := v.Min()
	petscgo.Printf("Max = %f\n", min)
	v.Scale(0.3)
	sum, _ = v.Sum()
	petscgo.Printf("Sum = %f\n", sum)

	err = v.Destroy()
	if err != nil {
		petscgo.Fatal(err)
	}

}
