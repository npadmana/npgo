package sf_test

import (
	. "github.com/npadmana/npgo/gsl/sf"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bessel", func() {
	Context("value at 0", func() {
		It("should be 1 for n=0", func() {
			Expect(BesselJ(0, 0)).To(BeNumerically("~", 1, 1.e-13))
		})
		It("should be 0 for n > 0", func() {
			for i := 1; i < 25; i++ {
				Expect(BesselJ(i, 0)).To(BeNumerically("~", 0, 1.e-13))
			}
		})
	})

	Context("test BesselJ at some values", func() {
		xx := []float64{0.1, 1.0, 2, 10.0, 100}
		j0x := []float64{0.99750156206604003228, 0.76519768655796655145, 0.22389077914123566805,
			-0.24593576445134833521, 0.019985850304223122424}
		j5x := []float64{2.6030817909644408340e-9, 0.00024975773021123443138, 0.0070396297558716854842,
			-0.23406152818679364044, -0.074195736964513920834}
		j19x := []float64{1.5677657562983752611e-42, 1.5484784412116534205e-23,
			7.8192432733637439506e-18, 0.000043146277524562556633, -0.038093921164499174989}
		It("should agree with the above", func() {
			for i, x1 := range xx {
				Expect(BesselJ(0, x1)).To(BeNumerically("~", j0x[i], 1.e-13))
				Expect(BesselJ(5, x1)).To(BeNumerically("~", j5x[i], 1.e-13))
				Expect(BesselJ(19, x1)).To(BeNumerically("~", j19x[i], 1.e-13))
			}
		})
	})

	Context("test BesselJarr at some values", func() {
		// J[0-9, 12.75]
		x := 12.75
		js := []float64{0.18288505664015526694, -0.12117855082319186611,
			-0.20189345676928340280, 0.057839427130867661307,
			0.22911201071322112576, 0.085917128610761280349,
			-0.16172602748909463151, -0.23812986036520328635,
			-0.099749897617795251674, 0.11295351825659748032,
			0.25921368809769757684}
		arr := BesselJArr(0, 10, x)
		It("should have length 11", func() {
			Expect(arr).To(HaveLen(11))
		})
		It("should agree with the above", func() {
			for i, j1 := range js {
				Expect(arr[i]).To(BeNumerically("~", j1, 1.e-13))
			}
		})
	})
})
