package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/npadmana/npgo/gsl/spline"
)

// The weight file is assumed to have a first header line, and then two columns
// with z and n(z)
func readWeight(wfn string, Pk float64) (*spline.Spline, error) {
	fmt.Println("Reading in ", wfn)
	ff, err := os.Open(wfn)
	if err != nil {
		return nil, err
	}
	defer ff.Close()
	buf := bufio.NewReader(ff)

	zz := make([]float64, 0)
	fkp := make([]float64, 0)

	// Skip first line
	var nitem int
	var z1, nz float64
	var char []byte
	err = nil
	for err == nil {
		char, err = buf.Peek(1)
		if err != nil {
			continue
		}
		if char[0] == '#' {
			buf.ReadString('\n')
		}
		nitem, err = fmt.Fscanln(buf, &z1, &nz)
		if err == nil {
			if nitem != 2 {
				return nil, errors.New("Unexpected number of elements")
			}
			zz = append(zz, z1)
			fkp = append(fkp, 1/(1+nz*Pk))
			fmt.Println(z1, 1/(1+nz*Pk))
		}
	}

	sp, err := spline.New(spline.Cubic, zz, fkp)
	if err != nil {
		return nil, err
	}

	// Return the spline
	return sp, nil

}

func main() {
	var wfn string
	var help bool
	var Pk float64
	var err error
	flag.StringVar(&wfn, "weight", "", "Spline weights")
	flag.BoolVar(&help, "help", false, "help")
	flag.Float64Var(&Pk, "Pk", 20000, "P0 in FKP weight")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if wfn == "" {
		log.Fatal("Need to specify a weight file")
	}

	sp, err := readWeight(wfn, Pk)
	if err != nil {
		log.Fatal(err)
	}
	defer sp.Free()

	var tmp float64
	for z1 := 0.43; z1 < 0.70; z1 = z1 + 0.01 {
		tmp, err = sp.Eval(z1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(z1, tmp)
	}
}
