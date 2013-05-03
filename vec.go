// Package petscgo contains bindings between PETSc and Go. 
package petscgo

/*
#cgo pkg-config: petsc ompi

#include "petsc.h"
#include "mypetsc.h"

*/
import "C"

import (
	"errors"
)

// Define a wrapper type; nothing is exported
type Vec struct {
	v C.Vec
}

// NewVec creates a new MPI vector of local size n
func NewVec(local, global int64) (*Vec, error) {
	v := new(Vec)
	perr := C.VecCreateMPI(C.PETSC_COMM_WORLD, C.PetscInt(local), C.PetscInt(global), &v.v)
	if perr != 0 {
		return nil, errors.New("Error creating vector")
	}
	return v, nil
}

// Destroy destroys the vector
func (v *Vec) Destroy() error {
	perr := C.VecDestroy(&v.v)
	if perr != 0 {
		return errors.New("Error destroying vector")
	}
	return nil
}

// Duplicate duplicates the vector
func (v *Vec) Duplicate() (*Vec, error) {
	v1 := new(Vec)
	perr := C.VecDuplicate(v.v, &v1.v)
	if perr != 0 {
		return nil, errors.New("Error destroying vector")
	}
	return v1, nil
}

// AssemblyBegin starts assembling the vector
func (v *Vec) AssemblyBegin() error {
	if perr := C.VecAssemblyBegin(v.v); perr != 0 {
		return errors.New("Error in AssemblyBegin")
	}
	return nil
}

// AssemblyEnd ends the vector assembly
func (v *Vec) AssemblyEnd() error {
	if perr := C.VecAssemblyEnd(v.v); perr != 0 {
		return errors.New("Error in AssemblyEnd")
	}
	return nil
}

// Copy copies src to dest
func Copy(src, dst *Vec) error {
	if perr := C.VecCopy(src.v, dst.v); perr != 0 {
		return errors.New("Error in Copy")
	}
	return nil
}

// LocalSize returns the local size 
func (v *Vec) LocalSize() (int64, error) {
	var ll C.PetscInt
	if perr := C.VecGetLocalSize(v.v, &ll); perr != 0 {
		return -1, errors.New("Error in AssemblyEnd")
	}
	return int64(ll), nil
}

// OwnRange returns the ownership range
func (v *Vec) OwnRange() (int64, int64, error) {
	var clo, chi C.PetscInt
	perr := C.VecGetOwnershipRange(v.v, &clo, &chi)
	if perr != 0 {
		return -1, -1, errors.New("Error getting Ownership Range")
	}
	return int64(clo), int64(chi), nil
}
