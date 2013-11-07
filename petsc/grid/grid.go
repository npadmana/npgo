// Package grid defines the interfaces that a grid must satisfy
package grid

// Gridder is a generic interface that all grids must satisfy
//
// The Array method returns an interface to allow different types to be returned.
type Gridder interface {
	Ndim() int           // Number of dimensions
	Dimensions() []int64 // Array of dimensions
	Strides() []int64    // Strides for each dimension
	Array() interface{}  // Returns the array
	Lo() []int64         // The lo indices in each dimension
	Hi() []int64         // The hi indices in each dimension
}
