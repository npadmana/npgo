package particles

// Define a simple particle interface
type Particle interface {
	Length() int64
	GetPos(ipart int64, idim int) float32
}
