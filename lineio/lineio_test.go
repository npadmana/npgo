package lineio

import (
	"testing"

	"github.com/npadmana/npgo/nptest"
)

var (
	eps = nptest.NewEps(1.e-9, 1.e-9)
)

func TestParseFloat64s(t *testing.T) {
	inp := []byte(" 1.234   7.9 -2.3 1.2e4  1.2e-4  ")
	var x, y, z, w, v float64
	err := ParseToFloat64s(inp, []byte{' '}, &x, &y, &z, &w, &v)
	if err != nil {
		t.Error(err)
	}
	eps.EqFloat64(1.234, x, "", t)
	eps.EqFloat64(7.9, y, "", t)
	eps.EqFloat64(-2.3, z, "", t)
	eps.EqFloat64(1.2e4, w, "", t)
	eps.EqFloat64(1.2e-4, v, "", t)
}

func TestParseFloat64Arr1(t *testing.T) {
	inp := []byte(" 1.234   7.9 -2.3 1.2e4  1.2e-4  ")
	truth := []float64{1.234, 7.9, -2.3, 1.2e4, 1.2e-4}
	var out []float64
	err := ParseToFloat64Arr(inp, []byte{' '}, &out, true)
	if err != nil {
		t.Error(err)
	}
	if len(out) != 5 {
		t.Error("Did not get the correct number of elements ", len(out))
	}
	for i := range truth {
		eps.EqFloat64(truth[i], out[i], "", t)
	}
}

func TestParseFloat64Arr2(t *testing.T) {
	inp := []byte(" 1.234   7.9 -2.3 1.2e4  1.2e-4  ")
	//truth := []float64{1.234, 7.9, -2.3, 1.2e4, 1.2e-4}
	out := make([]float64, 6)
	err := ParseToFloat64Arr(inp, []byte{' '}, &out, false)
	if err == nil {
		t.Error("An error was expected")
	}
}

func TestParseFloat64Arr3(t *testing.T) {
	inp := []byte(" 1.234   7.9 -2.3 1.2e4  1.2e-4  ")
	truth := []float64{1.234, 7.9, -2.3, 1.2e4, 1.2e-4}
	out := make([]float64, 5)
	err := ParseToFloat64Arr(inp, []byte{' '}, &out, false)
	if err != nil {
		t.Error(err)
	}
	if len(out) != 5 {
		t.Error("Did not get the correct number of elements ", len(out))
	}
	for i := range truth {
		eps.EqFloat64(truth[i], out[i], "", t)
	}
}

func TestParseFloat64Arr4(t *testing.T) {
	inp := []byte(" 1.234   7.9 -2.3 1.2e4  1.2e-4  ")
	truth := []float64{1.234, 7.9, -2.3, 1.2e4, 1.2e-4}
	out := make([]float64, 2)
	err := ParseToFloat64Arr(inp, []byte{' '}, &out, true)
	if err != nil {
		t.Error(err)
	}
	if len(out) != 5 {
		t.Error("Did not get the correct number of elements ", len(out))
	}
	for i := range truth {
		eps.EqFloat64(truth[i], out[i], "", t)
	}
}
