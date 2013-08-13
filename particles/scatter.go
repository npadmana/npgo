package particles

import (
	_ "unsafe"

	"github.com/npadmana/mpi"
	"github.com/npadmana/petscgo"
)

// Scatter takes the particles and reshuffles them across MPI ranks according to localid and
//
// localid is the local index of the particle, while mpirank is the destination rank for the particle.
// Note that this is completely general, and a particle may be shunted to many ranks (useful for ghosts).
// Also, note that localndx and mpirank will be modified by this routine.
//
func (p Particle) Scatter(localndx, mpirank []int64) error {

	// Get the rank and size
	rank, size := petscgo.RankSize()

	// Work out the final number of particles
	var npart_local, npart_final int64
	npart_local = int64(len(mpirank))
	mpi.AllReduceInt64(petscgo.WORLD, &npart_local, &npart_final, 1, mpi.SUM)
	petscgo.Printf("%d total particles expected after scatter\n", npart_final)

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

	_ = rank
	_ = icount
	_ = icount_check
	_ = narr1

	return nil
}
