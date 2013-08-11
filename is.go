package petscgo

/*
#cgo pkg-config: PETSc ompi

#include "petsc.h"
#include "petscis.h"

*/
import "C"

import (
	"errors"
	"unsafe"
)

type IS struct {
	is C.IS
}

// NewStrideIS creates a new index set, based on strides
func NewStrideIS(nlocal, first, step int64) (*IS, error) {
	ret := new(IS)
	perr := C.ISCreateStride(C.PETSC_COMM_WORLD, C.PetscInt(nlocal), C.PetscInt(first), C.PetscInt(step), &ret.is)
	if perr != 0 {
		return nil, errors.New("Error creating strided index set")
	}
	return ret, nil
}

// NewGeneralIS creates a general index set
//
// Note that this uses PetscCopyMode PETSC_COPY_VALUES; you are free to delete the indices after completion
func NewGeneralIS(nlocal int64, idx []int64) (*IS, error) {
	ret := new(IS)
	perr := C.ISCreateGeneral(C.PETSC_COMM_WORLD, C.PetscInt(nlocal), (*C.PetscInt)(unsafe.Pointer(&idx[0])), C.PETSC_COPY_VALUES, &ret.is)
	if perr != 0 {
		return nil, errors.New("Error creating general index set")
	}
	return ret, nil
}

// Destroy frees the index set
func (i *IS) Destroy() error {
	perr := C.ISDestroy(&i.is)
	if perr != 0 {
		return errors.New("Error destroying index set")
	}
	return nil
}
