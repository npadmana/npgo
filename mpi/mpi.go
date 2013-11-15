// Package mpi wraps any MPI commands I might need.
//
package mpi

/*
#cgo pkg-config: ompi mpich

#include <stdlib.h>
#include "mpi.h"

MPI_Op mpiop(int i) {
        MPI_Op retval;
        retval = MPI_SUM; // This prevents uninitialized warnings
        switch(i) {
        case 0 :
                retval = MPI_SUM;
                break;
        default :
                MPI_Abort(MPI_COMM_WORLD,1);
        }
        return retval;
}

MPI_Datatype mpitype(int i) {
        MPI_Datatype retval;
        retval = MPI_LONG; // This prevents uninitialized warnings
        switch(i) {
        case 0 :
                retval = MPI_LONG;
                break;
        default :
                MPI_Abort(MPI_COMM_WORLD,1);
        }
        return retval;
}


MPI_Comm retworld() {
	return MPI_COMM_WORLD;
}


*/
import "C"

import (
	"errors"
	"os"
	"unsafe"
)

type Comm C.MPI_Comm        // MPI Communicators
type Op C.MPI_Op            // MPI Optypes
type MpiType C.MPI_Datatype // MPI Datatypes

var (
	SUM = Op(C.mpiop(0)) // MPI_SUM
)

var (
	MPI_i64 = MpiType(C.mpitype(0))
	WORLD   = Comm(C.retworld)
)

// Initialize initializes the MPI environment
func Initialize() error {
	// Allocate space for argc and argv
	argc := C.int(len(os.Args))
	argv := make([](*C.char), argc)
	// Copy os.Args into argv
	for i, gstr := range os.Args {
		argv[i] = C.CString(gstr)
	}
	ptrargv := &argv[0]

	perr := C.MPI_Init(&argc, &ptrargv)
	if perr != 0 {
		return errors.New("Error initializing MPI")
	}

	// update os.Args
	os.Args = os.Args[0:0]
	for i := 0; i < int(argc); i++ {
		os.Args = append(os.Args, C.GoString(argv[i]))
		C.free(unsafe.Pointer(argv[i]))
	}

	return nil
}

// Finalize finalizes the MPI environment
func Finalize() error {
	perr := C.MPI_Finalize()
	if perr != 0 {
		return errors.New("Error initializing MPI")
	}
	return nil
}

// AllReduceInt64 : MPI_Allreduce for int64
func AllReduceInt64(comm Comm, in, out *int64, n int, op Op) {
	C.MPI_Allreduce(unsafe.Pointer(in), unsafe.Pointer(out), C.int(n), C.MPI_Datatype(MPI_i64), C.MPI_Op(SUM), C.MPI_Comm(comm))
}

// AllGatherInt64 : MPI_Allgather for int64
func AllGatherInt64(comm Comm, in, out []int64) {
	C.MPI_Allgather(unsafe.Pointer(&in[0]), C.int(len(in)), C.MPI_Datatype(MPI_i64), unsafe.Pointer(&out[0]), C.int(len(in)), C.MPI_Datatype(MPI_i64), C.MPI_Comm(comm))
}

// Abort calls MPI_Abort
func Abort(comm Comm, err int) error {
	perr := C.MPI_Abort(C.MPI_Comm(comm), C.int(err))
	if perr != 0 {
		return errors.New("Error aborting!?!!")
	}
	return nil
}

// Barrier calls MPI_Barrier
func Barrier(comm Comm) error {
	perr := C.MPI_Barrier(C.MPI_Comm(comm))
	if perr != 0 {
		return errors.New("Error calling Barrier")
	}
	return nil
}

// Rank returns the MPI_Rank
func Rank(comm Comm) (int, error) {
	var r C.int
	perr := C.MPI_Comm_rank(C.MPI_Comm(comm), &r)
	if perr != 0 {
		return -1, errors.New("Error calling MPI_Comm_rank")
	}
	return int(r), nil
}

// Size returns the MPI_Size
func Size(comm Comm) (int, error) {
	var r C.int
	perr := C.MPI_Comm_size(C.MPI_Comm(comm), &r)
	if perr != 0 {
		return -1, errors.New("Error calling MPI_Comm_size")
	}
	return int(r), nil
}

// TypeSize returns the size of an MPI type
func TypeSize(t1 MpiType) int {
	var n C.int
	C.MPI_Type_size(C.MPI_Datatype(t1), &n)
	return int(n)
}
