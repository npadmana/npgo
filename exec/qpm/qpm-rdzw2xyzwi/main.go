package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"

	"github.com/npadmana/npgo/cosmo"
	"github.com/npadmana/npgo/gnuplot"
	"github.com/npadmana/npgo/gsl/spline"
)

type Pos [3]float64

var (
	loPos = Pos{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64}
	hiPos = Pos{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}
)

func (p *Pos) Min(p1 Pos) {
	for i := range p {
		if p1[i] < p[i] {
			p[i] = p1[i]
		}
	}
}

func (p *Pos) Max(p1 Pos) {
	for i := range p {
		if p1[i] > p[i] {
			p[i] = p1[i]
		}
	}
}

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
		}
	}

	sp, err := spline.New(spline.Cubic, zz, fkp)
	if err != nil {
		return nil, err
	}

	// Test the spline
	plot, err := gnuplot.New(false)
	defer close(plot)
	if err != nil {
		return nil, err
	}
	plot <- "set term pngcairo"
	plot <- "set output 'fkp_test.png'"
	plot <- "plot '-' w points ps 3, '-' w lines lw 2"
	for i := range zz {
		plot <- fmt.Sprintln(zz[i], fkp[i])
	}
	plot <- "e"
	for z1 = zz[0]; z1 < zz[len(zz)-1]; z1 = z1 + 0.001 {
		nz, err = sp.Eval(z1)
		if err != nil {
			return nil, err
		}
		plot <- fmt.Sprint(z1, nz)
	}
	plot <- "e"
	plot <- "set output"

	// Return the spline
	return sp, nil

}

func splineCosmo(om, zmin, zmax float64) (*spline.Spline, error) {
	lcdm := cosmo.NewFlatLCDMSimple(om, 1)
	nz := 1000
	dz := (zmax - zmin) / float64(nz)
	avals := make([]float64, nz+1)
	zvals := make([]float64, nz+1)
	for i := 0; i <= nz; i++ {
		zvals[i] = zmin + float64(i)*dz
		avals[i] = cosmo.Z2A(zvals[i])
	}
	dist := cosmo.ComDis(lcdm, avals)

	sp, err := spline.New(spline.Cubic, zvals, dist)
	if err != nil {
		return nil, err
	}

	return sp, nil
}

func doOne(infn, outfn string, zmin, zmax float64, dist, fkp *spline.Spline, minpos, maxpos *Pos) error {
	var err error
	var ra, dec, zred, r, w, theta, phi float64
	var ncol, ind int
	var p Pos

	ff, err := os.Open(infn)
	if err != nil {
		return err
	}
	defer ff.Close()

	gg, err := os.Create(outfn)
	if err != nil {
		return err
	}
	defer gg.Close()

	ind = 1
	elim := 0
	maxz := -1.0
	minz := 10.0
	for err == nil {
		ncol, err = fmt.Fscanln(ff, &ra, &dec, &zred, &w)
		if err != nil {
			continue
		}
		if ncol != 4 {
			return errors.New("Error, incorrect number of columns read in " + infn)
		}
		if zred < minz {
			minz = zred
		}
		if zred > maxz {
			maxz = zred
		}
		if (zred < zmin) || (zred >= zmax) {
			elim++
			continue
		}
		theta = (math.Pi / 180) * (90 - dec)
		phi = (math.Pi / 180) * ra
		w, err = fkp.Eval(zred)
		if err != nil {
			panic("Error in fkp spline in " + infn)
		}
		r, err = dist.Eval(zred)
		if err != nil {
			panic("Error in dist spline " + infn)
		}
		p[0] = r * math.Sin(theta) * math.Cos(phi)
		p[1] = r * math.Sin(theta) * math.Sin(phi)
		p[2] = r * math.Cos(theta)
		_, err = fmt.Fprintf(gg, "%10.4f %10.4f %10.4f %7.4f %8d\n", p[0], p[1], p[2], w, ind)
		if err != nil {
			panic("Error while writing file")
		}
		minpos.Min(p)
		maxpos.Max(p)
		ind++
	}

	fmt.Printf("%s had zmin, zmax = %f, %f, %d objects removed \n", infn, minz, maxz, elim)
	return nil

}

func main() {
	tstart := time.Now()

	var wfn, infmt, outfmt string
	var help bool
	var Pk, zmin, zmax, om float64
	var nstart, nend int
	var err error
	flag.StringVar(&wfn, "weight", "", "Spline weights")
	flag.StringVar(&infmt, "in", "", "Input format eg. a%03i.rdzw")
	flag.StringVar(&outfmt, "out", "", "Output format eg. a%03i.dat")
	flag.BoolVar(&help, "help", false, "help")
	flag.Float64Var(&Pk, "Pk", 20000, "P0 in FKP weight")
	flag.Float64Var(&zmin, "zmin", 0.43, "minimum redshift, inclusive")
	flag.Float64Var(&zmax, "zmax", 0.7, "Maximum redshift, exclusive")
	flag.Float64Var(&om, "om", 0.274, "Omega_matter at z=0")
	flag.IntVar(&nstart, "nstart", 0, "starting index to fill in")
	flag.IntVar(&nend, "nend", 0, "ending index (exclusive)")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if wfn == "" {
		log.Fatal("Need to specify a weight file")
	}
	if infmt == "" {
		log.Fatal("Need to specify an input format")
	}
	if outfmt == "" {
		log.Fatal("Need to specify an output format")
	}

	fkp, err := readWeight(wfn, Pk)
	if err != nil {
		log.Fatal(err)
	}
	defer fkp.Free()

	dist, err := splineCosmo(om, zmin, zmax)
	if err != nil {
		log.Fatal(err)
	}
	defer dist.Free()

	var infn, outfn string
	minpos := hiPos
	maxpos := loPos
	var lock sync.Mutex
	var wg sync.WaitGroup

	for ii := nstart; ii < nend; ii++ {
		wg.Add(1)
		go func(ii int) {
			myminpos := hiPos
			mymaxpos := loPos
			infn = fmt.Sprintf(infmt, ii)
			outfn = fmt.Sprintf(outfmt, ii)
			fmt.Printf("Processing %s --> %s ...\n", infn, outfn)
			err = doOne(infn, outfn, zmin, zmax, dist, fkp, &myminpos, &mymaxpos)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("My min:", myminpos)
			fmt.Println("My max:", mymaxpos)
			lock.Lock()
			minpos.Min(myminpos)
			maxpos.Max(mymaxpos)
			lock.Unlock()
			wg.Done()
		}(ii)
	}

	wg.Wait()
	fmt.Println("The minimum particle position was at ", minpos)
	fmt.Println("The maximum particle position was at ", maxpos)

	fmt.Println("The total elapsed time was ", time.Since(tstart))

}
