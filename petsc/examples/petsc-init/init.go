package main

import (
	"github.com/npadmana/petscgo"
)

func main() {
	if err := petscgo.Initialize(); err != nil {
		petscgo.Fatal(err)
	}
	defer func() {
		if err := petscgo.Finalize(); err != nil {
			petscgo.Fatal(err)
		}
	}()
	rank, size := petscgo.RankSize()

	petscgo.Printf("Initialization successful\n")
	petscgo.SyncPrintf("Hello from rank %d of %d\n", rank, size)
	petscgo.SyncFlush()
}
