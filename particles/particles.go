// Package particles uses PETSc vectors to handle distributed particles
package particles

import (
	"errors"
	"fmt"
	"io"

	"github.com/npadmana/mpi"
	"github.com/npadmana/petscgo"
)

// Define a particle class
type Particle map[string](*petscgo.Vec)

// Define LocalParticle
type LocalParticle map[string]([]float64)

// New returns a Particle class. local/global defines the size of the array, while
// fieldnames is the list of field names
func New(fieldnames []string, local, global int64) (Particle, error) {
	pp := make(Particle, len(fieldnames))
	var err error
	for _, ff := range fieldnames {
		pp[ff], err = petscgo.NewVec(local, global)
		if err != nil {
			return nil, err
		}
	}
	return pp, nil
}

// Destroy deallocates the Particle data
func (pp Particle) Destroy() error {
	for _, v := range pp {
		if err := v.Destroy(); err != nil {
			return err
		}
	}
	return nil
}

// NewLocal makes a local particle class based on Particle
func (pp Particle) NewLocal(n int64) LocalParticle {
	lpp := make(LocalParticle, len(pp))
	for key := range pp {
		lpp[key] = make([]float64, n)
	}
	return lpp
}

// AssemblyBegin runs assembly begin on all the vectors
func (pp Particle) AssemblyBegin() error {
	for _, v := range pp {
		if err := v.AssemblyBegin(); err != nil {
			return err
		}
	}
	return nil
}

// AssemblyEnd runs assembly End on all the vectors
func (pp Particle) AssemblyEnd() error {
	for _, v := range pp {
		if err := v.AssemblyEnd(); err != nil {
			return err
		}
	}
	return nil
}

// SetValues sets LocalParticle into Particle
func (pp Particle) SetValues(ix []int64, lpp LocalParticle, add bool) error {
	for ff, v := range pp {
		if err := v.SetValues(ix, lpp[ff], add); err != nil {
			return err
		}
	}
	return nil
}

// GetArray gives you access to the Particle data
func (pp Particle) GetArray(fieldnames []string) (LocalParticle, error) {
	lpp := make(LocalParticle, len(fieldnames))
	for _, ff := range fieldnames {
		v, ok := pp[ff]
		if !ok {
			return nil, errors.New("Field name missing")
		}
		if err := v.GetArray(); err != nil {
			return nil, err
		}
		lpp[ff] = v.Arr
	}
	return lpp, nil
}

// RestoreArrays restores individual arrays
func (pp Particle) RestoreArray(lpp LocalParticle) error {
	for k := range lpp {
		delete(lpp, k)
		if err := pp[k].RestoreArray(); err != nil {
			return err
		}
	}
	return nil
}

// Dump writes out particles to an io.Writer instance
func (pp Particle) Dump(w io.Writer, fieldnames []string, format []string, hdr bool) error {
	nfields := len(fieldnames)

	var fmts []string
	switch len(format) {
	case 1:
		fmts = make([]string, nfields)
		for i := range fmts {
			fmts[i] = format[0]
		}
	case nfields:
		fmts = format
	default:
		return errors.New("format must have length 1 or len(fieldnames)")
	}

	// Get rank and size
	rank, size := petscgo.RankSize()

	// Get size of array
	v, ok := pp[fieldnames[0]]
	if !ok {
		return errors.New("Error accessing field when getting size")
	}
	npart, err := v.Size()
	if err != nil {
		return err
	}

	lpp1 := make([][]float64, nfields)
	if lpp, err := pp.GetArray(fieldnames); err == nil {
		for i, fns := range fieldnames {
			lpp1[i] = lpp[fns]
		}

		nlocal := len(lpp1[0])

		for irank := 0; irank < size; irank++ {
			if irank == rank {
				if hdr {
					if irank == 0 {
						fmt.Fprint(w, "# ")
						for _, fns := range fieldnames {
							fmt.Fprint(w, fns, " ")
						}
						fmt.Fprint(w, "\n")
					}
					fmt.Fprintf(w, "# Rank %d/%d with %d/%d particles \n", rank, size, nlocal, npart)
				}

				for ip := 0; ip < nlocal; ip++ {
					for i := range fieldnames {
						fmt.Fprintf(w, fmts[i], lpp1[i][ip])
						fmt.Fprint(w, " ")
					}
					fmt.Fprint(w, "\n")
				}

			}

			if err = mpi.Barrier(petscgo.WORLD); err != nil {
				return err
			}
		}

		err = pp.RestoreArray(lpp)
	}
	return err
}
