// Package petscgo contains bindings between PETSc and Go. 
package petscgo

/* 
Author : Nikhil Padmanabhan, Yale Univ.

This uses pkg-config to handle all the setup.
An example file is in the pkgconfig directory.

This assumes that you are using OpenMPI; if not, you will need to change the ompi
setting below

*/

/*
#cgo pkg-config: petsc ompi

#include "petsc.h"
#include "mypetsc.h"

*/
import "C"

import (
	"errors"
	"fmt"
	"log"
	"os"
	"unsafe"
)

// Some common PETSc variables
var (
	NULL     = C.PETSC_NULL
	NULLCHAR = C.petscNullChar()
)

// Fatal calls MPI_Abort
func Fatal(err error) {
	log.Println("Calling MPI_Abort : ", err)
	C.petscAbort()
}

// Initialize initializes the PETSc environment
func Initialize() error {
	if int(C.sizePetscReal()) != 8 {
		return errors.New("PetscScalar has the wrong size, expecting 64 bits")
	}
	if int(C.sizePetscInt()) != 8 {
		return errors.New("PetscInt has the wrong size, expecting 64 bits")
	}

	// Allocate space for argc and argv
	argc := C.int(len(os.Args))
	argv := make([](*C.char), argc)
	// Copy os.Args into argv
	for i, gstr := range os.Args {
		argv[i] = C.CString(gstr)
	}
	ptrargv := &argv[0]

	perr := C.PetscInitialize(&argc, &ptrargv, NULLCHAR, NULLCHAR)
	if perr != 0 {
		return errors.New("Error initializing PETSc")
	}

	// update os.Args
	os.Args = os.Args[0:0]
	for i := 0; i < int(argc); i++ {
		os.Args = append(os.Args, C.GoString(argv[i]))
		C.free(unsafe.Pointer(argv[i]))
	}

	return nil
}

// Finalize shuts it down
func Finalize() error {
	perr := C.PetscFinalize()
	if perr != 0 {
		return errors.New("Error finalizing PETSc")
	}
	return nil
}

// Print prints from the rank-0 process
func Printf(format string, args ...interface{}) error {
	cstr := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cstr))
	perr := C.mypetscPrintf(cstr)
	if perr != 0 {
		return errors.New("Error in Print")
	}
	return nil
}

// SyncPrint prints from all processes in order
func SyncPrintf(format string, args ...interface{}) error {
	cstr := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cstr))
	perr := C.mypetscSyncPrintf(cstr)
	if perr != 0 {
		return errors.New("Error in SyncPrint")
	}
	return nil
}

// SyncFlush flushes the SyncPrint buffer
func SyncFlush() error {
	if perr := C.PetscSynchronizedFlush(C.PETSC_COMM_WORLD); perr != 0 {
		return errors.New("Error in SyncFlush")
	}
	return nil
}

// RankSize returns the rank and size 
func RankSize() (int, int) {
	var cr, cs C.int
	C.petscRankSize(&cr, &cs)
	return int(cr), int(cs)
}
