package sf_test

import (
	. "github.com/npadmana/npgo/gsl/sf"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BesselJ", func() {
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

var _ = Describe("SphBessel", func() {
	Context("value at 0", func() {
		It("should be 1 for n=0", func() {
			Expect(SphBessel(0, 0)).To(BeNumerically("~", 1, 1.e-13))
		})
		It("should be 0 for n > 0", func() {
			for i := 1; i < 25; i++ {
				Expect(SphBessel(i, 0)).To(BeNumerically("~", 0, 1.e-13))
			}
		})
	})

	Context("test SphBessel at some values", func() {
		xx := []float64{0.1, 1.0, 2, 10.0, 100}
		j0x := []float64{0.99833416646828152307, 0.84147098480789650665,
			0.45464871341284084770, -0.054402111088936981340,
			-0.0050636564110975879366}
		j5x := []float64{9.6163102329164460441e-10, 0.000092561158611258163567,
			0.0026351697702441173490, -0.055534511621452180909,
			-0.0092901489349075717663}
		j7x := []float64{4.9318874757319734185e-14,
			4.7901341987394885770e-7, 0.000056096557033489486513,
			0.11338623065577473648, 0.0097006298438983563051}
		It("should agree with the above", func() {
			for i, x1 := range xx {
				Expect(SphBessel(0, x1)).To(BeNumerically("~", j0x[i], 1.e-13))
				Expect(SphBessel(5, x1)).To(BeNumerically("~", j5x[i], 1.e-13))
				Expect(SphBessel(7, x1)).To(BeNumerically("~", j7x[i], 1.e-13))
			}
		})
	})

	Context("test SphBesselArr at some values", func() {
		// J[0-9, 12.75]
		x := 12.75
		js := []float64{0.014321500755383059466, -0.075989485983702641022,
			-0.032201379810371916177, 0.063361493901203850364,
			0.066988082344366186965, -0.016075788716945365447,
			-0.080857390257024933625, -0.066367040564727115896,
			0.0027785190044047972768, 0.070071732570600178932,
			0.10164210208119546937}
		arr := SphBesselArr(10, x)
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
