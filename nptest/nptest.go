package nptest

import (
	"math"
	"testing"
)

type Eps struct {
	EpsAbs float64
	EpsRel float64
}

func NewEps(abs, rel float64) (e Eps) {
	e.EpsAbs = abs
	e.EpsRel = rel
	return
}

func (e Eps) EqFloat64(xtrue, xin float64, s string, t *testing.T) {
	dxt := xtrue * e.EpsRel
	dx := math.Abs(xtrue - xin)
	if (dx > dxt) && (dx > e.EpsAbs) {
		t.Errorf("%s : %f expected, %s got", s, xtrue, xin)
	}
}

func (e Eps) NeqFloat64(xtrue, xin float64, s string, t *testing.T) {
	dxt := xtrue * e.EpsRel
	dx := math.Abs(xtrue - xin)
	if (dx <= dxt) || (dx <= e.EpsAbs) {
		t.Errorf("%s : %f expected, %s got", s, xtrue, xin)
	}
}
