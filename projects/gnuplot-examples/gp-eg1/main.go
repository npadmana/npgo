package main

import (
	"github.com/npadmana/npgo/gnuplot"
	"log"
)

func main() {
	pp, err := gnuplot.New(false)
	if err != nil {
		log.Fatal(err)
	}

	pp <- "set term png"
	pp <- "set output \"sin.png\""
	pp <- "plot sin(x)"
	pp <- "set term pdfcairo"
	pp <- "set output \"cos.pdf\""
	pp <- "plot cos(x) lw 5"
	pp <- "set output"
	close(pp)

}
