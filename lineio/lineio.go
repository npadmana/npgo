package lineio

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// LineIOType is an interface that can be passed to the file reader.
// It has one method, parsing a single structure from an io.Reader and appending it
// on to the data. This may safely assume that comments etc have been parsed out from
// the file.
type LineIOType interface {
	Add(r io.Reader) error
}

// Parameters for line IO -- comment characters
// The default is "#"
type LineIOParams struct {
	Comment string
}

var (
	LineIOParamsDefault = LineIOParams{Comment: "#"}
)

// Parse in a file, based on the Line. Note that arr may be modified even
// if an error occured. The reader only has a single line in it.
func (l LineIOParams) Parse(fn string, arr LineIOType) error {
	// Open the file
	ff, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer ff.Close()

	scan := bufio.NewScanner(ff)
	scan.Split(bufio.ScanLines)
	var bb []byte
	var n int
	for scan.Scan() {
		// Trim out leading whitespace
		bb = scan.Bytes()
		bb = bytes.TrimSpace(bb)
		if n = bytes.IndexAny(bb, l.Comment); n != -1 {
			bb = bb[0:n]
		}
		if len(bb) > 0 {
			err = arr.Add(bytes.NewBuffer(bb))
			if err != nil {
				return err
			}
		}
	}
	return scan.Err()
}

// Read is equivalent to LineIOParamsDefault.Parse(fn,arr)
func Read(fn string, arr LineIOType) error {
	return LineIOParamsDefault.Parse(fn, arr)
}
