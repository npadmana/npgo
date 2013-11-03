// Count the number of lines, removing
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/npadmana/npgo/lineio"
)

type LC int

func (l *LC) Add(b []byte) error {
	*l++
	return nil
}

func main() {
	t := time.Now()
	var l LC
	fn := os.Args[1]
	if err := lineio.Read(fn, &l); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s has %d non-skipped lines (elapsed time:%s)\n", fn, l, time.Since(t))
}
