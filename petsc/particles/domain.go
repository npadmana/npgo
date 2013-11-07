package particles

type Domainer interface {
	Domain(p Particle) ([]int64, []int64)
}

type Slab struct {
	L       float32
	N, Idim int
}

func (s Slab) Domain(p Particle) ([]int64, []int64) {
	ll := p.Length()
	dx := s.L / float32(s.N)
	localndx := make([]int64, ll)
	mpirank := make([]int64, ll)

	for ipart := int64(0); ipart < ll; ipart++ {
		x := p.GetPos(ipart, s.Idim)
		ix := int64(x / dx)
		localndx[ipart] = ipart
		mpirank[ipart] = ix
	}

	return localndx, mpirank
}
