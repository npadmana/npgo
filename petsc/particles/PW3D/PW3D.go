package PW3D

import (
	"math/rand"

	"github.com/npadmana/petscgo"
	"github.com/npadmana/petscgo/particles"
	"github.com/npadmana/petscgo/structvec"
)

// Define an example particle type
type One struct {
	Pos [3]float32
	W   float32
}

type Arr []One

func (p Arr) Length() int64 {
	return int64(len(p))
}

func (p Arr) GetPos(ipart int64, idim int) float32 {
	return p[ipart].Pos[idim]
}

func NewVec(nlocal, nglobal int64) *structvec.StructVec {
	v, err := structvec.NewStructVec(One{}, nlocal, nglobal)
	if err != nil {
		petscgo.Fatal(err)
	}
	return v
}

func GetArray(s *structvec.StructVec) Arr {
	return Arr(s.GetArray().([]One))
}

func (p Arr) FillRandom(Lmax, Wmax float32) {
	for i := range p {
		for idim := range p[i].Pos {
			p[i].Pos[idim] = rand.Float32() * Lmax
		}
		p[i].W = rand.Float32() * Wmax
	}
}

func DomainDecompose(d particles.Domainer, s *structvec.StructVec) {
	pp := GetArray(s)
	localndx, mpirank := d.Domain(pp)
	s.RestoreArray()
	s.Scatter(localndx, mpirank)
}
