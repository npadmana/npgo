package nptest

import (
	"testing"
)

var eps = NewEps(1.e-8, 1.e-5)

func TestEqFloat64(t *testing.T) {
	eps.EqFloat64(1., 1., "1", t)
	eps.NeqFloat64(1., 1.01, "2", t)
	eps.EqFloat64(1., 1+1.e-6, "3", t)
	eps.NeqFloat64(1., 1+1.e-3, "4", t)
	eps.EqFloat64(0, 1.e-10, "5", t)
	eps.NeqFloat64(0, 1.e-6, "6", t)
}
