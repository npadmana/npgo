package petscgo

/*
#cgo pkg-config: PETSc ompi

#include "petsc.h"
#include "mypetsc.h"

*/
import "C"

import (
	"errors"
	"reflect"
	"unsafe"
)

// Different types of vector norms
type Norms int

const (
	NORM1 Norms = iota
	NORM2
	NORMINF
)

// Define a wrapper type; nothing is exported
type Vec struct {
	v   C.Vec
	ptr *C.PetscScalar
	Arr []float64 // Access local vector storage using Get/RestoreArray
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

// Copy copies v (src) to dest
func (v *Vec) Copy(dst *Vec) error {
	if perr := C.VecCopy(v.v, dst.v); perr != 0 {
		return errors.New("Error in Copy")
	}
	return nil
}

// LocalSize returns the local size
func (v *Vec) LocalSize() (int64, error) {
	var ll C.PetscInt
	if perr := C.VecGetLocalSize(v.v, &ll); perr != 0 {
		return -1, errors.New("Error in LocalSize")
	}
	return int64(ll), nil
}

// Size returns the global size
func (v *Vec) Size() (int64, error) {
	var ll C.PetscInt
	if perr := C.VecGetSize(v.v, &ll); perr != 0 {
		return -1, errors.New("Error in GlobalSize")
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

// Range returns the ownership ranges of all processors
func (v *Vec) Ranges() ([]int64, error) {
	_, size := RankSize()
	var ptr *C.PetscInt
	perr := C.VecGetOwnershipRanges(v.v, &ptr)
	if perr != 0 {
		return nil, errors.New("Error getting ownership ranges")
	}
	var rr []int64
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&rr)))
	sliceHeader.Cap = size + 1
	sliceHeader.Len = size + 1
	sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
	return rr, nil
}

// Set sets the vector to a value
func (v *Vec) Set(a float64) error {
	perr := C.VecSet(v.v, C.PetscScalar(a))
	if perr != 0 {
		return errors.New("Error in Set")
	}
	return nil
}

// GetArray sets the Arr
func (v *Vec) GetArray() error {
	size, err := v.LocalSize()
	if err != nil {
		return err
	}
	perr := C.VecGetArray(v.v, &v.ptr)
	if perr != 0 {
		return errors.New("Error getting array")
	}
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&v.Arr)))
	sliceHeader.Cap = int(size)
	sliceHeader.Len = int(size)
	sliceHeader.Data = uintptr(unsafe.Pointer(v.ptr))
	return nil
}

// RestoreArray undoes GetArray; this resets the slice as well to prevent accidents
func (v *Vec) RestoreArray() error {
	perr := C.VecRestoreArray(v.v, &v.ptr)
	if perr != 0 {
		return errors.New("Error restoring array")
	}
	v.Arr = v.Arr[0:0]
	return nil
}

//SetValues sets values based on global indices. If add is true, then use ADD_VALUES, otherwise INSERT_VALUES.
//Must be followed by AssemblyBegin/End
func (v *Vec) SetValues(ix []int64, y []float64, add bool) error {
	var iora C.InsertMode = C.INSERT_VALUES
	if add {
		iora = C.ADD_VALUES
	}
	perr := C.VecSetValues(v.v, C.PetscInt(len(ix)), (*C.PetscInt)(unsafe.Pointer(&ix[0])), (*C.PetscScalar)(unsafe.Pointer(&y[0])), iora)
	if perr != 0 {
		return errors.New("Error setting values")
	}
	return nil
}

// Sum returns the sum of all components of the array
func (v *Vec) Sum() (float64, error) {
	var sum C.PetscScalar
	perr := C.VecSum(v.v, &sum)
	if perr != 0 {
		return 0, errors.New("Error in sum")
	}
	return float64(sum), nil
}

//SqrtAbs takes the sqrt of the absolute value
func (v *Vec) SqrtAbs() error {
	perr := C.VecSqrtAbs(v.v)
	if perr != 0 {
		return errors.New("Error in SqrtAbs")
	}
	return nil
}

//Abs takes the absolute value of components
func (v *Vec) Abs() error {
	perr := C.VecAbs(v.v)
	if perr != 0 {
		return errors.New("Error in Abs")
	}
	return nil
}

// Dot takes the dot product
func (v *Vec) Dot(v1 *Vec) (float64, error) {
	var ret C.PetscScalar
	perr := C.VecDot(v.v, v1.v, &ret)
	if perr != 0 {
		return 0, errors.New("Error in Dot")
	}
	return float64(ret), nil
}

// Reciprocal takes the reciprocal
func (v *Vec) Reciprocal() error {
	perr := C.VecReciprocal(v.v)
	if perr != 0 {
		return errors.New("Error in Reciprocal")
	}
	return nil
}

// Scale scales the vector
func (v *Vec) Scale(a float64) error {
	perr := C.VecScale(v.v, C.PetscScalar(a))
	if perr != 0 {
		return errors.New("Error in Scale")
	}
	return nil
}

// Shift adds a constant to the vector
func (v *Vec) Shift(a float64) error {
	perr := C.VecShift(v.v, C.PetscScalar(a))
	if perr != 0 {
		return errors.New("Error in Shift")
	}
	return nil
}

// Max returns the maximum and its location
func (v *Vec) Max() (float64, int64, error) {
	var val C.PetscReal
	var p C.PetscInt
	perr := C.VecMax(v.v, &p, &val)
	if perr != 0 {
		return 0, -1, errors.New("Error in Max")
	}
	return float64(val), int64(p), nil
}

// Min returns the minimum and its location
func (v *Vec) Min() (float64, int64, error) {
	var val C.PetscReal
	var p C.PetscInt
	perr := C.VecMin(v.v, &p, &val)
	if perr != 0 {
		return 0, -1, errors.New("Error in Min")
	}
	return float64(val), int64(p), nil
}

// Norm returns an appropriate norm of the vector
func (v *Vec) Norm(n Norms) (float64, error) {
	var n1 C.NormType
	switch n {
	case NORM1:
		n1 = C.NORM_1
	case NORM2:
		n1 = C.NORM_2
	case NORMINF:
		n1 = C.NORM_INFINITY
	default:
		return -1, errors.New("Unknown norm type")
	}
	var val C.PetscReal
	perr := C.VecNorm(v.v, n1, &val)
	if perr != 0 {
		return -1, errors.New("Error in Norm")
	}
	return float64(val), nil
}

// AXPBY sets y = alpha x + beta y
func (y *Vec) AXPBY(x *Vec, alpha, beta float64) error {
	perr := C.VecAXPBY(y.v, C.PetscScalar(alpha), C.PetscScalar(beta), x.v)
	if perr != 0 {
		return errors.New("Error in AXPBY")
	}
	return nil
}

// AXPY sets y = alpha x + y
func (y *Vec) AXPY(x *Vec, alpha float64) error {
	perr := C.VecAXPY(y.v, C.PetscScalar(alpha), x.v)
	if perr != 0 {
		return errors.New("Error in AXPY")
	}
	return nil
}

// AYPX sets y = x + alpha y
func (y *Vec) AYPX(x *Vec, alpha float64) error {
	perr := C.VecAYPX(y.v, C.PetscScalar(alpha), x.v)
	if perr != 0 {
		return errors.New("Error in AYPX")
	}
	return nil
}

// WAXPY sets w = alpha x +  y
func (w *Vec) WAXPY(x, y *Vec, alpha float64) error {
	perr := C.VecWAXPY(w.v, C.PetscScalar(alpha), x.v, y.v)
	if perr != 0 {
		return errors.New("Error in AYPX")
	}
	return nil
}
