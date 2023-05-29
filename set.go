// package set implements the set data structure, and its common methods
// leveraging the fact that all comparable types are hashable in Go
// and empty structs occupy no memory, we can use a map[T]struct{} to
// implement a set.
// The methods are implemented as described in https://doc.rust-lang.org/std/collections/struct.HashSet.html (inc)
// The underlying map is only accessible via the methods of the set.
package set

import "fmt"

// Set is a set data structure
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet returns a new set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

// Add adds an element to the set
func (s *Set[T]) Add(e T) {
	s.m[e] = struct{}{}
}

// Contains returns true if the set contains the element
func (s *Set[T]) Contains(e T) bool {
	_, ok := s.m[e]
	return ok
}

// Remove removes an element from the set
func (s *Set[T]) Remove(e T) {
	delete(s.m, e)
}

// Intersection returns the intersection of two sets as a new set IE the values that are in both sets
func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	// make sure s is the smaller set
	if len(s.m) > len(s2.m) {
		s, s2 = s2, s
	}
	s3 := NewSet[T]()
	for k := range s.m {
		if ok := s2.Contains(k); ok {
			s3.Add(k)
		}
	}
	return s3
}

// Union returns the union of two sets as a new set IE the values that are in either set without duplicates
func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for k := range s.m {
		s3.Add(k)
	}
	for k := range s2.m {
		s3.Add(k)
	}
	return s3
}

// Difference returns the values in s that are not in s2 as a new set
func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for k := range s.m {
		if ok := s2.Contains(k); !ok {
			s3.Add(k)
		}
	}
	return s3
}

// SymmetricDifference returns the values that are in one of the sets, but not both
func (s *Set[T]) SymmetricDifference(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for k := range s.m {
		if ok := s2.Contains(k); !ok {
			s3.Add(k)
		}
	}
	for k := range s2.m {
		if ok := s.Contains(k); !ok {
			s3.Add(k)
		}
	}
	return s3
}

// IsSubset returns true if s is a subset of s2 IE all values in s are in s2
func (s *Set[T]) IsSubset(s2 *Set[T]) bool {
	for k := range s.m {
		if ok := s2.Contains(k); !ok {
			return false
		}
	}
	return true
}

// IsDisjoint returns true if s and s2 have no common values IE their intersection is empty
func (s *Set[T]) IsDisjoint(s2 *Set[T]) bool {
	if len(s.m) > len(s2.m) {
		s, s2 = s2, s
	}
	for k := range s.m {
		if ok := s2.Contains(k); ok {
			return false
		}
	}
	return true
}

// IsSuperset returns true if s is a superset of s2 IE all values in s2 are in s
func (s *Set[T]) IsSuperset(s2 *Set[T]) bool {
	return s2.IsSubset(s)
}

// IsEqual returns true if s and s2 contain the same values
func (s *Set[T]) IsEqual(s2 *Set[T]) bool {
	return s.IsSubset(s2) && s.IsSuperset(s2)
}

// Copy returns a copy of the set
func (s *Set[T]) Copy() *Set[T] {
	s2 := NewSet[T]()
	for k := range s.m {
		s2.Add(k)
	}
	return s2
}

// Len returns the number of elements in the set
func (s *Set[T]) Len() int {
	return len(s.m)
}

// Clear removes all elements from the set
func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

// IsEmpty returns true if the set is empty
func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

// ToSlice returns a slice of the elements in the set
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, s.Len())
	for k := range s.m {
		slice = append(slice, k)
	}
	return slice
}

// Filter returns a new set containing only the elements that satisfy the predicate
func (s *Set[T]) Filter(predicate func(T) bool) *Set[T] {
	s2 := NewSet[T]()
	for k := range s.m {
		if predicate(k) {
			s2.Add(k)
		}
	}
	return s2
}

// Map returns a new set containing the results of applying the function to each element
func (s *Set[T]) Map(f func(T) T) *Set[T] {
	s2 := NewSet[T]()
	for k := range s.m {
		s2.Add(f(k))
	}
	return s2
}

// ForEach applies the function to each element in the set
func (s *Set[T]) ForEach(f func(T)) {
	for k := range s.m {
		f(k)
	}
}

// Reduce applies the function to each element in the set and returns the result
func (s *Set[T]) Reduce(f func(T, T) T) T {
	var result T
	for k := range s.m {
		result = f(result, k)
	}
	return result
}

// Any returns true if any element in the set satisfies the predicate
func (s *Set[T]) Any(predicate func(T) bool) bool {
	for k := range s.m {
		if predicate(k) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the set satisfy the predicate
func (s *Set[T]) All(predicate func(T) bool) bool {
	for k := range s.m {
		if !predicate(k) {
			return false
		}
	}
	return true
}

// String returns a string representation of the set
func (s *Set[T]) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}
