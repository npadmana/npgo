package main

import (
	"github.com/npadmana/petscgo"
)

func main() {
	petscgo.Initialize()
	defer petscgo.Finalize()
	rank, size := petscgo.RankSize()

	// Set up basic elements for the vector
	var bs, nlocal, nlocal1 int64
	bs = 2                // block size
	nlocal1 = 5           // number of local blocks
	nlocal = nlocal1 * bs // local size
	r64 := int64(rank) * nlocal1
	s64 := int64(size) * nlocal1
	gndx := []int64{(r64 + nlocal1) % s64, (s64 + r64 - 1) % s64}
	petscgo.SyncPrintf("Ghost indices : %v \n", gndx)
	petscgo.SyncFlush()

	// Create the vector
	v, _ := petscgo.NewGhostVecBlocked(nlocal, petscgo.DETERMINE, bs, gndx)
	defer v.Destroy()

	// Fill in the local versions of the array
	lo, _, _ := v.OwnRange()
	v.GetArray()
	for ii := range v.Arr {
		v.Arr[ii] = float64(int64(ii) + lo)
	}
	v.RestoreArray()

	petscgo.Printf("Filled in vector\n")

	// Update ghost values
	v.GhostUpdateBegin(false, true)
	v.GhostUpdateEnd(false, true)

	// Get the local values and print them
	lv, _ := v.GhostGetLocalForm()
	lv.GetArray()
	petscgo.SyncPrintf("Rank %d : ", rank)
	for _, val := range lv.Arr {
		petscgo.SyncPrintf("%3d ", int(val))
	}
	petscgo.SyncPrintf("\n")
	petscgo.SyncFlush()
	lv.RestoreArray()
	lv.Destroy()

	// Now reset the array to 0
	v.Set(0)
	v.GhostUpdateBegin(false, true)
	v.GhostUpdateEnd(false, true)
	// Fill the array with 1's including the ghosts
	lv, _ = v.GhostGetLocalForm()
	lv.GetArray()
	for ii := range lv.Arr {
		lv.Arr[ii] = float64(rank + 1)
	}
	lv.RestoreArray()
	lv.Destroy()
	v.GhostUpdateBegin(true, false)
	v.GhostUpdateEnd(true, false)

	// Reprint, only with local pieces
	v.GetArray()
	petscgo.SyncPrintf("Rank %d : ", rank)
	for _, val := range v.Arr {
		petscgo.SyncPrintf("%3d ", int(val))
	}
	petscgo.SyncPrintf("\n")
	petscgo.SyncFlush()
	v.RestoreArray()

}
