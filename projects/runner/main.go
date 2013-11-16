package main

import (
	"bytes"
	"fmt"
	"github.com/npadmana/npgo/lineio"
	"github.com/npadmana/npgo/mpi"
	"log"
	"os"
	"os/exec"
	"unicode"
)

type LogFile struct {
	F          *os.File
	linenum    int
	Rank, Size int
}

func (l *LogFile) Add(b []byte) error {
	// Extract the first word --- which will be the command; the rest will be args
	if (l.linenum % l.Size) == l.Rank {
		ndx := bytes.IndexFunc(b, unicode.IsSpace)
		var comm, args string
		if ndx == -1 {
			comm = string(b)
			args = ""
		} else {
			comm = string(b[0:ndx])
			args = string(b[ndx:])
		}
		fmt.Fprintln(l.F, "-->", string(b))
		out, err := exec.Command(comm, args).CombinedOutput()
		if err != nil {
			fmt.Fprintln(l.F, "-->ERROR : ", err)
			fmt.Fprintln(l.F, "-->Output follows :")
		}
		fmt.Fprintln(l.F, string(out))
		fmt.Fprintln(l.F, "-->")
	}
	l.linenum += 1

	return nil
}

func main() {
	var err error

	mpi.Initialize()
	defer mpi.Finalize()

	var rank, size int
	if rank, err = mpi.Rank(mpi.WORLD); err != nil {
		mpi.Abort(mpi.WORLD, -1)
	}
	if size, err = mpi.Size(mpi.WORLD); err != nil {
		mpi.Abort(mpi.WORLD, -1)
	}

	// Define the input and output strings
	if len(os.Args) != 3 {
		fmt.Println("Incorrect number of parameters")
		mpi.Abort(mpi.WORLD, 1)
	}
	infn := os.Args[1]
	outfn := os.Args[2]

	// Open output file
	var outff LogFile
	outff.F, err = os.Create(fmt.Sprintf("%s-%05d.out", outfn, rank))
	if err != nil {
		log.Printf("Unable to open file %s-%05d.out", outfn, rank)
	}
	defer outff.F.Close()
	outff.Rank, outff.Size, outff.linenum = rank, size, 0

	lineio.Read(infn, &outff)

	mpi.Barrier(mpi.WORLD)

}
