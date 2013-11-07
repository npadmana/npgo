package main

import (
	"fmt"
	"math/rand"

	"github.com/npadmana/npgo/petsc"
	"github.com/npadmana/npgo/petsc/structvec"
)

type pstruct struct {
	x, y, z, w float32
}

func (p *pstruct) FillRandom() {
	p.x = rand.Float32()
	p.y = rand.Float32()
	p.z = rand.Float32()
	p.w = rand.Float32()
}

func (p pstruct) String() string {
	return fmt.Sprintf("(%7.4f, %7.4f, %7.4f, %7.4f)", p.x, p.y, p.z, p.w)
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
	rank, _ := petsc.RankSize()

	v, err := structvec.NewStructVec(pstruct{}, petsc.DECIDE, 10)
	if err != nil {
		petsc.Fatal(err)
	}
	defer v.Destroy()
	petsc.Printf("Type of v : %s\n", v.Type())
	petsc.Printf("Size of v : %d\n", v.BlockSize())
	petsc.SyncPrintf("Local size = %d\n", v.Nlocal)
	petsc.SyncFlush()
	petsc.Printf("Global size = %d\n", v.Ntotal)

	// local particle data
	lpp, ok := v.GetArray().([]pstruct)
	if !ok {
		petsc.Fatal(err)
	}
	for i := range lpp {
		lpp[i].FillRandom()
	}
	err = v.RestoreArray()
	if err != nil {
		petsc.Fatal(err)
	}

	// Print array
	lpp, ok = v.GetArray().([]pstruct)
	if !ok {
		petsc.Fatal(err)
	}
	for i := range lpp {
		petsc.SyncPrintf("%s\n", lpp[i])
	}
	petsc.SyncFlush()
	err = v.RestoreArray()
	if err != nil {
		petsc.Fatal(err)
	}

	petsc.Printf("----------------\n")

	// Fiddle with array
	if rank == 0 {
		lpp = make([]pstruct, 2)
		ix := []int64{3, 7}
		err = v.SetValues(ix, lpp)
		if err != nil {
			petsc.Fatal(err)
		}
	}
	v.AssemblyBegin()
	v.AssemblyEnd()

	// Print array
	lpp, ok = v.GetArray().([]pstruct)
	if !ok {
		petsc.Fatal(err)
	}
	for i := range lpp {
		petsc.SyncPrintf("%s\n", lpp[i])
	}
	petsc.SyncFlush()
	err = v.RestoreArray()
	if err != nil {
		petsc.Fatal(err)
	}

}
