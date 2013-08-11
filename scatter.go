package petscgo

/*
#cgo pkg-config: PETSc ompi

#include "petsc.h"

*/
import "C"

import (
	"errors"
)

// Scatter wraps the VecScatter context
type Scatter struct {
	s C.VecScatter
}

// NewScatter creates a new scatter context
func NewScatter(src, dest *Vec, sis, dis *IS) (*Scatter, error) {
	s := new(Scatter)
	perr := C.VecScatterCreate(src.v, sis.is, dest.v, dis.is, &s.s)
	if perr != 0 {
		return nil, errors.New("Error creating scatter context")
	}
	return s, nil
}

// Destroy destroys the scatter context
func (s *Scatter) Destroy() error {
	perr := C.VecScatterDestroy(&s.s)
	if perr != 0 {
		return errors.New("Error destroying scatter context")
	}
	return nil
}

// Begin starts a scatter
func (s *Scatter) Begin(src, dest *Vec, add, forward bool) error {
	var iora C.InsertMode = C.INSERT_VALUES
	if add {
		iora = C.ADD_VALUES
	}
	var mode C.ScatterMode = C.SCATTER_FORWARD
	if !forward {
		mode = C.SCATTER_REVERSE
	}
	perr := C.VecScatterBegin(s.s, src.v, dest.v, iora, mode)
	if perr != 0 {
		return errors.New("Error starting scatter")
	}
	return nil
}

// End ends the scatter
func (s *Scatter) End(src, dest *Vec, add, forward bool) error {
	var iora C.InsertMode = C.INSERT_VALUES
	if add {
		iora = C.ADD_VALUES
	}
	var mode C.ScatterMode = C.SCATTER_FORWARD
	if !forward {
		mode = C.SCATTER_REVERSE
	}
	perr := C.VecScatterEnd(s.s, src.v, dest.v, iora, mode)
	if perr != 0 {
		return errors.New("Error ending scatter")
	}
	return nil
}
