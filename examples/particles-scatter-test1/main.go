package main

import (
	"os"

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
	var np1 int64 = 1
	if rank == 0 {
		np1 = 2
	}
	pp, err := particles.New([]string{"x", "y", "z"}, np1, petscgo.DETERMINE)
	if err != nil {
		petscgo.Fatal(err)
	}
	defer pp.Destroy()
	n1, _ := pp["x"].LocalSize()
	petscgo.SyncPrintf("%d/%d rank has local size %d\n", rank, size, n1)
	petscgo.SyncFlush()

	// Fill these in
	lpp, err := pp.GetArray([]string{"x", "y", "z"})
	if err != nil {
		petscgo.Fatal(err)
	}
	map1 := map[string]float64{"x": 1, "y": 2, "z": 3}
	for k, v := range lpp {
		for i := range v {
			v[i] = (float64(i) + 1) * (map1[k] + float64(rank)*10)
		}
	}
	err = pp.RestoreArray(lpp)
	if err != nil {
		petscgo.Fatal(err)
	}

	err = pp.Dump(os.Stdout, []string{"x", "y", "z"}, []string{"%6.2f"}, true)
	if err != nil {
		petscgo.Fatal(err)
	}

	// Set up scatters
	var localndx, mpirank []int64
	if rank == 0 {
		localndx = make([]int64, 1+size)
		mpirank = make([]int64, 1+size)
		for i := 0; i < size; i++ {
			localndx[i] = 0
			mpirank[i] = int64(i)
		}
		localndx[size] = 1
		mpirank[size] = int64((rank + 1) % size)
	} else {
		localndx = make([]int64, 1)
		mpirank = make([]int64, 1)
		localndx[0] = 0
		mpirank[0] = int64((rank + 1) % size)
	}

	err = pp.Scatter(localndx, mpirank)
	if err != nil {
		petscgo.Fatal(err)
	}

}
