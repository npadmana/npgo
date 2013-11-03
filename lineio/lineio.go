package lineio

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strconv"
	"unsafe"
)

// The following is taken from Issue 2632 in go. We don't export this function, because that
// would be bad form.
func unsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// LineIOType is an interface that can be passed to the file reader.
// It has one method, parsing a single structure from an io.Reader and appending it
// on to the data. This may safely assume that comments etc have been parsed out from
// the file.
type LineIOType interface {
	Add([]byte) error
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
			err = arr.Add(bb)
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

// ParseLineToFloat64 parses a line into a sequence of float64s. This throws
// an error if the number of input arguments does not equal the number of elements parseable
func ParseToFloat64s(s, sep []byte, args ...*float64) error {
	barr := bytes.Split(s, sep)
	iarg := 0
	nargs := len(args)
	var err error
	var val float64
	for i := range barr {
		if len(barr[i]) == 0 {
			continue
		}
		if val, err = strconv.ParseFloat(unsafeString(barr[i]), 64); err != nil {
			return err
		}
		if (iarg + 1) > nargs {
			return errors.New("Number of elements does not equal number of parameters")
		}
		*args[iarg] = val
		iarg++
	}
	if iarg != len(args) {
		return errors.New("Number of elements does not equal number of parameters")
	}
	return nil
}

// ParseLineToFloat64 parses a line into a sequence of float64s. This throws
// an error if the number of input arguments does not equal the number of elements parseable.
// If grow is true, then the slice is automatically appended to.
func ParseToFloat64Arr(s, sep []byte, arr *[]float64, grow bool) error {
	barr := bytes.Split(s, sep)
	iarg := 0
	nargs := len(*arr)
	var err error
	var val float64
	for i := range barr {
		if len(barr[i]) == 0 {
			continue
		}
		if val, err = strconv.ParseFloat(unsafeString(barr[i]), 64); err != nil {
			return err
		}
		if (iarg + 1) > nargs {
			if grow {
				*arr = append(*arr, val)
			} else {
				return errors.New("Number of elements does not equal number of parameters")
			}
		} else {
			(*arr)[iarg] = val
		}
		iarg++
	}
	if iarg != len(*arr) {
		return errors.New("Number of elements does not equal number of parameters")
	}
	return nil
}
