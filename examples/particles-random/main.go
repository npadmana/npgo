package main

import (
	"math/rand"

	"github.com/npadmana/petscgo"
	"github.com/npadmana/petscgo/particles"
)

func main() {
	// PETSc initialization
	if err := petscgo.Initialize(); err != nil {
		petscgo.Fatal(err)
	}
	defer func() {
		if err := petscgo.Finalize(); err != nil {
			petscgo.Fatal(err)
		}
	}()
	rank, size := petscgo.RankSize()

	// Create particles
	pp, err := particles.New([]string{"x", "y", "z"}, petscgo.DECIDE, 100)
	if err != nil {
		petscgo.Fatal(err)
	}
	defer pp.Destroy()
	n1, _ := pp["x"].LocalSize()
	petscgo.SyncPrintf("%d/%d rank has local size %d\n", rank, size, n1)
	petscgo.SyncFlush()

	// Set z = 1
	pp["z"].Set(1)

	// Fill these in with random numbers
	lpp, err := pp.GetArray([]string{"x", "y"})
	if err != nil {
		petscgo.Fatal(err)
	}

	var tmp []float64
	tmp = lpp["x"]
	var i int64
	for i = 0; i < n1; i++ {
		tmp[i] = rand.Float64()
	}
	tmp = lpp["y"]
	for i = 0; i < n1; i++ {
		tmp[i] = rand.Float64() * 3
	}

	err = pp.RestoreArray(lpp)
	if err != nil {
		petscgo.Fatal(err)
	}

	sum, _ := pp["x"].Sum()
	petscgo.Printf("x Sum = %f\n", sum)
	sum, _ = pp["y"].Sum()
	petscgo.Printf("y Sum = %f\n", sum)
	sum, _ = pp["z"].Sum()
	petscgo.Printf("z Sum = %f\n", sum)
}
