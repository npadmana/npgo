package integ

import (
	"github.com/npadmana/npgo/gsl"
	"math"
	"testing"
)

var (
	cf    = func(x float64) float64 { return 1 }
	sinf  = func(x float64) float64 { return math.Sin(x) }
	ep    = func(x float64) float64 { return math.Exp(-x) }
	recip = func(x float64) float64 { return 1 / (x * x) }
	gauss = func(x float64) float64 { return math.Exp(-(x*x)/2) / (math.Sqrt2 * math.SqrtPi) }
)

func TestInteg1(t *testing.T) {
	w := NewWork(1000)
	defer w.Free()
	res, err := Qags(cf, gsl.Interval{0, 1}, gsl.Eps{1e-7, 1e-7}, w)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if math.Abs(res.Res-1) > 1e-7 {
		t.Errorf("Integration failed : expected=%f, actual=%f, error=%f", 1, res.Res, res.Err)
	}
}

func TestInteg2(t *testing.T) {
	w := NewWork(1000)
	defer w.Free()
	res, err := Qags(sinf, gsl.Interval{0, 2 * math.Pi}, gsl.Eps{1e-7, 1e-7}, w)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if math.Abs(res.Res) > 1e-7 {
		t.Errorf("Integration failed : expected=%f, actual=%f, error=%f", 1, res.Res, res.Err)
	}
}

func TestInteg3(t *testing.T) {
	w := NewWork(1000)
	defer w.Free()
	res, err := Qags(ep, gsl.Interval{1, gsl.Inf}, gsl.Eps{1e-7, 1e-7}, w)
	y0 := math.Exp(-1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if math.Abs(res.Res-y0) > 1e-7 {
		t.Errorf("Integration failed : expected=%f, actual=%f, error=%f", y0, res.Res, res.Err)
	}
}

func TestInteg4(t *testing.T) {
	w := NewWork(1000)
	defer w.Free()
	res, err := Qags(recip, gsl.Interval{gsl.NInf, -1}, gsl.Eps{1e-7, 1e-7}, w)
	y0 := 1.0
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if math.Abs(res.Res-y0) > 1e-7 {
		t.Errorf("Integration failed : expected=%f, actual=%f, error=%f", y0, res.Res, res.Err)
	}
}

func TestInteg5(t *testing.T) {
	w := NewWork(1000)
	defer w.Free()
	res, err := Qags(gauss, gsl.Interval{gsl.NInf, gsl.Inf}, gsl.Eps{1e-7, 1e-7}, w)
	y0 := 1.0
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if math.Abs(res.Res-y0) > 1e-7 {
		t.Errorf("Integration failed : expected=%f, actual=%f, error=%f", y0, res.Res, res.Err)
	}
}
