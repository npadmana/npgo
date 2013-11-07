package structvec

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/npadmana/petscgo"
)

type StructVec struct {
	v              *petscgo.Vec // Store the vector
	bs             int64        // Store the blocksize
	t              reflect.Type // Store the structure information
	Nlocal, Ntotal int64        // Local & global sizes
}

func Partition(nglobal int64) int64 {
	rank, size := petscgo.RankSize()
	retval := nglobal / int64(size)
	rem := nglobal % int64(size)

	if int64(rank) < rem {
		retval += 1
	}

	return retval
}

func NewStructVec(st interface{}, nlocal, nglobal int64) (*StructVec, error) {
	var err error
	r := new(StructVec)
	r.t = reflect.TypeOf(st)

	// Return if st is not a struct
	if r.t.Kind() != reflect.Struct {
		return nil, errors.New("")
	}

	// Block size
	r.bs = int64(r.t.Size())
	if (r.bs % 8) != 0 {
		return nil, errors.New("Struct size must be divisible by 8")
	}
	r.bs = r.bs / 8

	// Set nlocal if not set
	if nlocal == petscgo.DECIDE {
		nlocal = Partition(nglobal)
	}

	// Now allocate the vector
	r.Nlocal = nlocal
	r.v, err = petscgo.NewVecBlocked(r.Nlocal*r.bs, petscgo.DETERMINE, r.bs)
	if err != nil {
		return nil, errors.New("Error allocating vector")
	}
	gsize, _ := r.v.Size()
	r.Ntotal = int64(gsize) / r.bs

	return r, nil
}

func (s *StructVec) Type() reflect.Type {
	return s.t
}

func (s *StructVec) BlockSize() int64 {
	return s.bs
}

func (s *StructVec) Destroy() error {
	return s.v.Destroy()
}

func (s *StructVec) GetArray() interface{} {
	if err := s.v.GetArray(); err != nil {
		petscgo.Fatal(err)
	}

	// Make sure the lengths match
	if int64(len(s.v.Arr)) != (s.Nlocal * s.bs) {
		petscgo.Fatal(errors.New("Unexpected size of array"))
	}

	sliceHeaderIn := (*reflect.SliceHeader)(unsafe.Pointer(&s.v.Arr))
	ptr := reflect.New(reflect.SliceOf(s.t))
	sliceHeaderOut := (*reflect.SliceHeader)(unsafe.Pointer(ptr.Pointer()))
	sliceHeaderOut.Cap = int(s.Nlocal)
	sliceHeaderOut.Len = int(s.Nlocal)
	sliceHeaderOut.Data = sliceHeaderIn.Data

	return ptr.Elem().Interface()
}

func (s *StructVec) RestoreArray() error {
	if err := s.v.RestoreArray(); err != nil {
		return err
	}

	return nil
}

// AssemblyBegin runs assembly begin on all the vectors
func (s *StructVec) AssemblyBegin() error {
	if err := s.v.AssemblyBegin(); err != nil {
		return err
	}
	return nil
}

// AssemblyEnd runs assembly End on all the vectors
func (s *StructVec) AssemblyEnd() error {
	if err := s.v.AssemblyEnd(); err != nil {
		return err
	}
	return nil
}

// SetValues sets values
func (s *StructVec) SetValues(ix []int64, arr interface{}) error {
	t1 := reflect.TypeOf(arr)
	v1 := reflect.ValueOf(arr)
	if t1.Kind() != reflect.Slice {
		return errors.New("arr needs to be a slice")
	}
	if t1.Elem() != s.t {
		return errors.New("slice type does not match existing structure type")
	}
	ptr := v1.Index(0).Addr().Pointer()
	return s.v.SetValuesBlockedPtr(ix, ptr, false)
}

// OwnRange returns the ownership range
func (s *StructVec) OwnRange() (int64, int64, error) {
	lo, hi, err := s.v.OwnRange()
	if err != nil {
		return -1, -1, nil
	}
	return lo / s.bs, hi / s.bs, nil
}
