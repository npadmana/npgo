package cosmo

import (
	"math/rand"
	"testing"

	"github.com/npadmana/npgo/nptest"
)

func TestA2Z(t *testing.T) {
	eps := nptest.NewEps(1.e-8, 1.e-6)
	eps.EqFloat64(0, A2Z(1), "z=0, a=1", t)
	eps.EqFloat64(1, A2Z(0.5), "z=1,a=0.5", t)
	eps.EqFloat64(3, A2Z(0.25), "z=3,a=0.25", t)
}

func TestZ2A(t *testing.T) {
	eps := nptest.NewEps(1.e-8, 1.e-6)
	eps.EqFloat64(1, Z2A(0), "z=0,a=1", t)
	eps.EqFloat64(0.5, Z2A(1), "z=1,a=0.5", t)
	eps.EqFloat64(0.25, Z2A(3), "z=3,a=0.25", t)
}

func TestA2Z2A(t *testing.T) {
	eps := nptest.NewEps(1.e-8, 1.e-6)
	var a float64
	for i := 0; i < 1000; i++ {
		a = 1 - rand.Float64()
		eps.EqFloat64(a, A2Z(Z2A(a)), "A->Z->A", t)
	}
}

func TestZ2A2Z(t *testing.T) {
	eps := nptest.NewEps(1.e-8, 1.e-6)
	var z float64
	for i := 0; i < 1000; i++ {
		z = 100 * rand.Float64()
		eps.EqFloat64(z, Z2A(A2Z(z)), "Z->A->Z", t)
	}
}
