// Mostly testing speed
package main

import (
	"fmt"
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

func (arr *RDZWArr) Add(s []byte) error {
	var x RDZW
	if err := lineio.ParseToFloat64s(s, []byte{' '}, &x.ra, &x.dec, &x.z, &x.w); err != nil {
		return err
	}
	*arr = append(*arr, x)
	return nil
}

func main() {
	t := time.Now()
	var wg sync.WaitGroup
	for i := 1; i < len(os.Args); i++ {
		wg.Add(1)
		go func(fn string) {
			var l RDZWArr
			if err := lineio.Read(fn, &l); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s has %d records\n", fn, len(l))
			wg.Done()
		}(os.Args[i])
	}
	wg.Wait()
	fmt.Printf("Elapsed time : %s\n", time.Since(t))
}
