package cosmo

import (
	"fmt"
	"testing"

	"github.com/npadmana/npgo/nptest"
)

func TestComDis1(t *testing.T) {
	// These numbers from Ned Wright's cosmology calculator
	zvals := []float64{0.1, 0.3, 0.7, 1, 2}
	dists := []float64{413.5, 1185.3, 2505.2, 3317.1, 5244.3}
	avals := make([]float64, len(zvals))
	for i, z1 := range zvals {
		avals[i] = Z2A(z1)
	}
	lcdm := NewFlatLCDMSimple(0.27, 0.71)
	dists2 := ComDis(lcdm, avals)
	eps := nptest.NewEps(0.001, 0.001)
	for i, d1 := range dists {
		eps.EqFloat64(d1, dists2[i], fmt.Sprintf("Testing z=%f", zvals[i]), t)
	}
}

func BenchmarkComDis1(b *testing.B) {
	da := 0.1 / float64(b.N+1)
	lcdm := NewFlatLCDMSimple(0.27, 0.71)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ComDis(lcdm, []float64{1 - float64(i+1)*da})
	}
}

func BenchmarkComDis2(b *testing.B) {
	da := 0.1 / float64(b.N+1)
	avals := make([]float64, b.N)
	for i := range avals {
		avals[i] = 1 - float64(i+1)*da
	}
	lcdm := NewFlatLCDMSimple(0.27, 0.71)
	b.ResetTimer()
	ComDis(lcdm, avals)
}
