package main

import (
	"fmt"

	"github.com/npadmana/npgo/petsc"
	"github.com/npadmana/npgo/petsc/structvec"
)

type pstruct struct {
	pos [3]float32
	w   float32
}

func (p pstruct) String() string {
	return fmt.Sprintf("(%7.4f, %7.4f, %7.4f, %7.4f)", p.pos[0], p.pos[1], p.pos[2], p.w)
}

func dump(pp []pstruct, rank int) {
	petsc.SyncPrintf("--- rank : %d \n", rank)
	for i := range pp {
		petsc.SyncPrintf("%s\n", pp[i])
	}
	petsc.SyncFlush()
}

func main() {
	// PETSc initialization
	if err := petsc.Initialize(); err != nil {
		petsc.Fatal(err)
	}
	defer func() {
		if err := petsc.Finalize(); err != nil {
			petsc.Fatal(err)
		}
	}()
	rank, size := petsc.RankSize()

	// Create particles
	var np1 int64 = 1
	if rank == 0 {
		np1 = 2
	}
	pp, err := structvec.NewStructVec(pstruct{}, np1, petsc.DETERMINE)
	if err != nil {
		petsc.Fatal(err)
	}
	defer pp.Destroy()

	lpp, _ := pp.GetArray().([]pstruct)
	for i := range lpp {
		for j := 0; j < 3; j++ {
			lpp[i].pos[j] = (float32(i) + 1) * (float32(j + 1 + rank*10))
		}
	}
	pp.RestoreArray()

	lpp, _ = pp.GetArray().([]pstruct)
	dump(lpp, rank)
	pp.RestoreArray()

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

	petsc.Printf("\n\n\n")
	pp.Scatter(localndx, mpirank)
	lpp, _ = pp.GetArray().([]pstruct)
	dump(lpp, rank)
	pp.RestoreArray()

}
