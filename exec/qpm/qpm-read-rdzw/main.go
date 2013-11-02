// Mostly testing speed
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/npadmana/npgo/lineio"
)

type RDZW struct {
	ra, dec, z, w float64
}

type RDZWArr []RDZW

func (arr *RDZWArr) Add(r io.Reader) error {
	var x RDZW
	n, err := fmt.Fscan(r, &x.ra, &x.dec, &x.z, &x.w)
	if err != nil {
		return err
	}
	if n != 4 {
		return errors.New("Failed to parse line")
	}
	*arr = append(*arr, x)
	return nil
}

func main() {
	t := time.Now()
	var wg sync.WaitGroup
	for i := 1; i < len(os.Args)-1; i++ {
		wg.Add(1)
		go func(fn string) {
			var l RDZWArr
			if err := lineio.Read(fn, &l); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s has %d records\n", fn, len(l))
		}(os.Args[i])
	}
	fmt.Printf("Elapsed time : %s\n", time.Since(t))
}
