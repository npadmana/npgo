package structvec

import (
	"errors"

	"github.com/npadmana/npgo/mpi"
	"github.com/npadmana/npgo/petsc"
)

// Scatter takes the particles and reshuffles them across MPI ranks according to localid and
//
// localid is the local index of the particle, while mpirank is the destination rank for the particle.
// Note that this is completely general, and a particle may be shunted to many ranks (useful for ghosts).
// Also, note that localndx and mpirank will be modified by this routine.
//
func (s *StructVec) Scatter(localndx, mpirank []int64) {

	// Get the rank and size
	rank, size := petsc.RankSize()

	// Work out the final number of particles
	var npart_local, npart_final int64
	npart_local = int64(len(mpirank))
	mpi.AllReduceInt64(petsc.WORLD, &npart_local, &npart_final, 1, mpi.SUM)
	//petsc.Printf("%d total particles expected after scatter\n", npart_final)

	// Allocate arrays
	// narr[rank] are the number of particles headed for rank from the current rank (hereafter crank)
	// narr1[crank*size + rank] collects all the narr's, allowing every processor to know what index it needs to send objects to.
	// icount keeps track of indices that objects need to go to.
	// icount_check is used for assertion tests, to make sure nothing bad happened.
	narr := make([]int64, size)
	icount := make([]int64, size)
	narr1 := make([]int64, size*size)
	icount_check := make([]int64, size)

	// Loop over mpirank, incrementing narr
	for _, irank := range mpirank {
		narr[irank] += 1
	}
	mpi.AllGatherInt64(petsc.WORLD, narr, narr1)
	//petsc.Printf("%v\n", narr1)

	// Reset narr, icount
	for i := range narr {
		narr[i] = 0
		icount[i] = 0
	}

	for i := 0; i < size; i++ {
		// narr now holds the total number of local particles
		for j := 0; j < size; j++ {
			narr[i] += narr1[i+j*size]
		}
		// icount now holds the number of particles from ranks before my rank, on rank i
		for j := 0; j < rank; j++ {
			icount[i] += narr1[i+j*size]
		}
	}

	// Now switch icount to global indices
	// icount_check is the expected final global index
	rtot := int64(0)
	for i := 0; i < size; i++ {
		icount[i] += rtot
		rtot += narr[i]
		icount_check[i] = icount[i] + narr1[i+rank*size]
	}
	// Assertion check
	if rtot != npart_final {
		petsc.Fatal(errors.New("ASSERTION FAILURE : rtot != npart_final"))
	}

	// Now we start updating localndx and mpirank
	lo, _, _ := s.OwnRange()
	irank := int64(0)
	for i := range mpirank {
		localndx[i] += lo
		irank = mpirank[i]
		mpirank[i] = icount[irank]
		icount[irank]++
	}

	// Assertion check
	for i := range icount {
		if icount[i] != icount_check[i] {
			petsc.Fatal(errors.New("ASSERTION FAILURE : icount != icount_check"))
		}
	}

	// Create destination
	vecnew, err := petsc.NewVecBlocked(narr[rank]*s.bs, petsc.DETERMINE, s.bs)
	if err != nil {
		petsc.Fatal(err)
	}
	// Create index sets
	isin, err := petsc.NewBlockedIS(s.bs, npart_local, localndx)
	if err != nil {
		petsc.Fatal(err)
	}
	isout, err := petsc.NewBlockedIS(s.bs, npart_local, mpirank)
	if err != nil {
		petsc.Fatal(err)
	}
	defer isin.Destroy()
	defer isout.Destroy()

	// Create scatter context
	ctx, err := petsc.NewScatter(s.v, vecnew, isin, isout)
	defer ctx.Destroy()
	ctx.Begin(s.v, vecnew, false, true)
	ctx.End(s.v, vecnew, false, true)

	// Clean up
	s.v.Destroy()
	s.v = vecnew
	s.Nlocal = narr[rank]
	s.Ntotal = npart_final

}
