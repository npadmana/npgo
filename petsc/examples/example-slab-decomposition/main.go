package main

import (
	"github.com/npadmana/npgo/petsc"
	"github.com/npadmana/npgo/petsc/particles"
	"github.com/npadmana/npgo/petsc/particles/PW3D"
)

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

	pp := PW3D.NewVec(petsc.DECIDE, 10000)
	defer pp.Destroy()

	lpp := PW3D.GetArray(pp)
	lpp.FillRandom(1, 1)
	pp.RestoreArray()
	petsc.Printf("Generating random particles....\n")

	slab := particles.Slab{L: 1, N: size, Idim: 0}
	PW3D.DomainDecompose(slab, pp)
	petsc.Printf("Slab decomposition complete\n")

	lpp = PW3D.GetArray(pp)
	_, mpirank := slab.Domain(lpp)
	rank64 := int64(rank)
	petsc.SyncPrintf("# Rank %d has %d particles....\n", rank, lpp.Length())
	for ipart, irank := range mpirank {
		if irank != rank64 {
			petsc.SyncPrintf("ERROR: %d expected, %d placed, %+v\n", rank, irank, lpp[ipart])
		}
	}
	petsc.SyncFlush()
	pp.RestoreArray()

}
