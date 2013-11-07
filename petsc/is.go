package petsc

/*
#cgo pkg-config: PETSc ompi mpich

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
// icopy == true means that the indices are copied, whereas icopy == false means that the PETSC_USE_POINTER is called.
// note that if icopy==false, you must *NOT* delete this pointer.
func NewGeneralIS(nlocal int64, idx []int64, icopy bool) (*IS, error) {
	var mode C.PetscCopyMode
	switch icopy {
	case true:
		mode = C.PETSC_COPY_VALUES
	case false:
		mode = C.PETSC_USE_POINTER
	}

	ret := new(IS)
	perr := C.ISCreateGeneral(C.PETSC_COMM_WORLD, C.PetscInt(nlocal), (*C.PetscInt)(unsafe.Pointer(&idx[0])), mode, &ret.is)
	if perr != 0 {
		return nil, errors.New("Error creating general index set")
	}
	return ret, nil
}

// NewBlockedIS creates a blocked index set
//
// icopy == true means that the indices are copied, whereas icopy == false means that the PETSC_USE_POINTER is called.
// note that if icopy==false, you must *NOT* delete this pointer.
func NewBlockedIS(bs, nlocal int64, idx []int64) (*IS, error) {
	ret := new(IS)
	perr := C.ISCreateBlock(C.PETSC_COMM_WORLD, C.PetscInt(bs), C.PetscInt(nlocal), (*C.PetscInt)(unsafe.Pointer(&idx[0])), C.PETSC_COPY_VALUES, &ret.is)
	if perr != 0 {
		return nil, errors.New("Error creating blocked index set")
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
