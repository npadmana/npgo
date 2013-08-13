package main

import (
	"fmt"
	"math/rand"

	"github.com/npadmana/petscgo"
	"github.com/npadmana/petscgo/particles"
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
	if err := petscgo.Initialize(); err != nil {
		petscgo.Fatal(err)
	}
	defer func() {
		if err := petscgo.Finalize(); err != nil {
			petscgo.Fatal(err)
		}
	}()
	rank, _ := petscgo.RankSize()

	v, err := particles.NewStructVec(pstruct{}, petscgo.DECIDE, 10)
	if err != nil {
		petscgo.Fatal(err)
	}
	defer v.Destroy()
	petscgo.Printf("Type of v : %s\n", v.Type())
	petscgo.Printf("Size of v : %d\n", v.BlockSize())
	petscgo.SyncPrintf("Local size = %d\n", v.Nlocal)
	petscgo.SyncFlush()
	petscgo.Printf("Global size = %d\n", v.Ntotal)

	// local particle data
	lpp, ok := v.GetArray().([]pstruct)
	if !ok {
		petscgo.Fatal(err)
	}
	for i := range lpp {
		lpp[i].FillRandom()
	}
	err = v.RestoreArray()
	if err != nil {
		petscgo.Fatal(err)
	}

	// Print array
	lpp, ok = v.GetArray().([]pstruct)
	if !ok {
		petscgo.Fatal(err)
	}
	for i := range lpp {
		petscgo.SyncPrintf("%s\n", lpp[i])
	}
	petscgo.SyncFlush()
	err = v.RestoreArray()
	if err != nil {
		petscgo.Fatal(err)
	}

	petscgo.Printf("----------------\n")

	// Fiddle with array
	if rank == 0 {
		lpp = make([]pstruct, 2)
		ix := []int64{3, 7}
		err = v.SetValues(ix, lpp)
		if err != nil {
			petscgo.Fatal(err)
		}
	}
	v.AssemblyBegin()
	v.AssemblyEnd()

	// Print array
	lpp, ok = v.GetArray().([]pstruct)
	if !ok {
		petscgo.Fatal(err)
	}
	for i := range lpp {
		petscgo.SyncPrintf("%s\n", lpp[i])
	}
	petscgo.SyncFlush()
	err = v.RestoreArray()
	if err != nil {
		petscgo.Fatal(err)
	}

}
