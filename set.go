// package set implements the set data structure, and its common methods
// leveraging the fact that all comparable types are hashable in Go
// and empty structs occupy no memory, we can use a map[T]struct{} to
// implement a set
package set

// Set is a set data structure
type Set[T comparable] struct {
	M map[T]struct{}
}

// NewSet returns a new set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		M: make(map[T]struct{}),
	}
}

// Add adds an element to the set
func (s *Set[T]) Add(e T) {
	s.M[e] = struct{}{}
}

// Get returns the element if it exists in the set
func (s *Set[T]) Get(e T) (T, bool) {
	_, ok := s.M[e]
	return e, ok
}

// Delete deletes an element from the set
func (s *Set[T]) Delete(e T) {
	delete(s.M, e)
}

// Intersection returns the intersection of two sets
func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	// make sure s is the smaller set
	if len(s.M) > len(s2.M) {
		s, s2 = s2, s
	}
	s3 := NewSet[T]()
	for k := range s.M {
		if _, ok := s2.Get(k); ok {
			s3.Add(k)
		}
	}
	return s3
}

// Union returns the union of two sets
func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for k := range s.M {
		s3.Add(k)
	}
	for k := range s2.M {
		s3.Add(k)
	}
	return s3
}

// Difference returns the difference of two sets
func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	// make sure s is the smaller set
	if len(s.M) > len(s2.M) {
		s, s2 = s2, s
	}
	s3 := NewSet[T]()
	for k := range s.M {
		if _, ok := s2.Get(k); !ok {
			s3.Add(k)
		}
	}
	return s3
}
