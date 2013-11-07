package main

import (
	"github.com/npadmana/petscgo"
	"github.com/npadmana/petscgo/particles"
	"github.com/npadmana/petscgo/particles/PW3D"
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

	pp := PW3D.NewVec(petscgo.DECIDE, 10000)
	defer pp.Destroy()

	lpp := PW3D.GetArray(pp)
	lpp.FillRandom(1, 1)
	pp.RestoreArray()
	petscgo.Printf("Generating random particles....\n")

	slab := particles.Slab{L: 1, N: size, Idim: 0}
	PW3D.DomainDecompose(slab, pp)
	petscgo.Printf("Slab decomposition complete\n")

	lpp = PW3D.GetArray(pp)
	_, mpirank := slab.Domain(lpp)
	rank64 := int64(rank)
	petscgo.SyncPrintf("# Rank %d has %d particles....\n", rank, lpp.Length())
	for ipart, irank := range mpirank {
		if irank != rank64 {
			petscgo.SyncPrintf("ERROR: %d expected, %d placed, %+v\n", rank, irank, lpp[ipart])
		}
	}
	petscgo.SyncFlush()
	pp.RestoreArray()

}
