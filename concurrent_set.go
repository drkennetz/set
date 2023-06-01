package set

import (
	"fmt"
	"sync"
)

// ThreadSafeSet is a thread-safe set data structure
// We use a mutex lock to make sure that only one thread can access the set at a time
// We redefine a lot of code in this snippet because we don't want to lock and unlock the mutex
// for every operation. We only lock and unlock the mutex once per operation.
type ThreadSafeSet[T comparable] struct {
	m map[T]struct{}
	l sync.Mutex
}

// NewThreadSafeSet returns a new thread-safe set
func NewThreadSafeSet[T comparable]() *ThreadSafeSet[T] {
	return &ThreadSafeSet[T]{
		m: make(map[T]struct{}),
	}
}

// NewThreadSafeSetFromSlice returns a new thread-safe set from a slice
func NewThreadSafeSetFromSlice[T comparable](s []T) *ThreadSafeSet[T] {
	set := NewThreadSafeSet[T]()
	for _, v := range s {
		set.m[v] = struct{}{}
	}
	return set
}

// Add adds an element to the set
func (s *ThreadSafeSet[T]) Add(e T) {
	s.l.Lock()
	defer s.l.Unlock()
	s.m[e] = struct{}{}
}

// Contains returns true if the set contains the element
func (s *ThreadSafeSet[T]) Contains(e T) bool {
	s.l.Lock()
	defer s.l.Unlock()
	_, ok := s.m[e]
	return ok
}

// Remove removes an element from the set
func (s *ThreadSafeSet[T]) Remove(e T) {
	s.l.Lock()
	defer s.l.Unlock()
	delete(s.m, e)
}

// Pop removes and returns an arbitrary element from the set or returns the zero value of T if the set is empty
func (s *ThreadSafeSet[T]) Pop() T {
	s.l.Lock()
	defer s.l.Unlock()
	var zero T
	for k := range s.m {
		// subtle differences come in
		// we don't want to lock and lock again so we recode
		delete(s.m, k)
		return k
	}
	return zero
}

// Intersection returns the intersection of two sets as a new set IE the values that are in both sets
func (s *ThreadSafeSet[T]) Intersection(s2 *ThreadSafeSet[T]) *ThreadSafeSet[T] {
	s.l.Lock()
	defer s.l.Unlock()
	s2.l.Lock()
	defer s2.l.Unlock()
	// make sure s is the smaller set
	if len(s.m) > len(s2.m) {
		s, s2 = s2, s
	}
	// we don't lock s3 because it is created here
	s3 := NewThreadSafeSet[T]()
	for k := range s.m {
		if _, ok := s2.m[k]; ok {
			s3.m[k] = struct{}{}
		}
	}
	return s3
}

// Union returns the union of two sets as a new set IE all the values in both sets
func (s *ThreadSafeSet[T]) Union(s2 *ThreadSafeSet[T]) *ThreadSafeSet[T] {
	s.l.Lock()
	defer s.l.Unlock()
	s2.l.Lock()
	defer s2.l.Unlock()
	// we don't lock s3 because it is created here
	s3 := NewThreadSafeSet[T]()
	for k := range s.m {
		s3.m[k] = struct{}{}
	}
	for k := range s2.m {
		s3.m[k] = struct{}{}
	}
	return s3
}

// Difference returns the difference of two sets as a new set IE all the values in the first set that are not in the second set
func (s *ThreadSafeSet[T]) Difference(s2 *ThreadSafeSet[T]) *ThreadSafeSet[T] {
	s.l.Lock()
	defer s.l.Unlock()
	s2.l.Lock()
	defer s2.l.Unlock()
	// we don't lock s3 because it is created here
	s3 := NewThreadSafeSet[T]()
	for k := range s.m {
		if _, ok := s2.m[k]; !ok {
			s3.m[k] = struct{}{}
		}
	}
	return s3
}

// SymmetricDifference returns the symmetric difference of two sets as a new set IE all the values that are in one set but not both
func (s *ThreadSafeSet[T]) SymmetricDifference(s2 *ThreadSafeSet[T]) *ThreadSafeSet[T] {
	s.l.Lock()
	defer s.l.Unlock()
	s2.l.Lock()
	defer s2.l.Unlock()
	// we don't lock s3 because it is created here
	s3 := NewThreadSafeSet[T]()
	for k := range s.m {
		if _, ok := s2.m[k]; !ok {
			s3.m[k] = struct{}{}
		}
	}
	for k := range s2.m {
		if _, ok := s.m[k]; !ok {
			s3.m[k] = struct{}{}
		}
	}
	return s3
}

// IsSubset returns true if the first set is a subset of the second set
func (s *ThreadSafeSet[T]) IsSubset(s2 *ThreadSafeSet[T]) bool {
	s.l.Lock()
	defer s.l.Unlock()
	s2.l.Lock()
	defer s2.l.Unlock()
	for k := range s.m {
		if _, ok := s2.m[k]; !ok {
			return false
		}
	}
	return true
}

// IsSuperset returns true if the first set is a superset of the second set
func (s *ThreadSafeSet[T]) IsSuperset(s2 *ThreadSafeSet[T]) bool {
	// we can reuse IsSubset because it locks and unlocks the mutexes
	return s2.IsSubset(s)
}

// IsDisjoint returns true if the two sets have no elements in common
func (s *ThreadSafeSet[T]) IsDisjoint(s2 *ThreadSafeSet[T]) bool {
	s.l.Lock()
	defer s.l.Unlock()
	s2.l.Lock()
	defer s2.l.Unlock()
	// make sure s is the smaller set
	if len(s.m) > len(s2.m) {
		s, s2 = s2, s
	}
	for k := range s.m {
		if _, ok := s2.m[k]; ok {
			return false
		}
	}
	return true
}

// IsEqual returns true if the two sets contain the same values
func (s *ThreadSafeSet[T]) IsEqual(s2 *ThreadSafeSet[T]) bool {
	return s.IsSubset(s2) && s.IsSuperset(s2)
}

// Copy returns a copy of the set
func (s *ThreadSafeSet[T]) Copy() *ThreadSafeSet[T] {
	s.l.Lock()
	defer s.l.Unlock()
	// we don't lock s2 because it is created here
	s2 := NewThreadSafeSet[T]()
	for k := range s.m {
		s2.m[k] = struct{}{}
	}
	return s2
}

// Len returns the number of elements in the set
func (s *ThreadSafeSet[T]) Len() int {
	s.l.Lock()
	defer s.l.Unlock()
	return len(s.m)
}

// Clear removes all elements from the set
func (s *ThreadSafeSet[T]) Clear() {
	s.l.Lock()
	defer s.l.Unlock()
	s.m = make(map[T]struct{})
}

// IsEmpty returns true if the set is empty
func (s *ThreadSafeSet[T]) IsEmpty() bool {
	return s.Len() == 0
}

// ToSlice returns the set as a slice
func (s *ThreadSafeSet[T]) ToSlice() []T {
	s.l.Lock()
	defer s.l.Unlock()
	slice := make([]T, 0, len(s.m))
	for k := range s.m {
		slice = append(slice, k)
	}
	return slice
}

// Filter returns a new set containing only the elements that pass the predicate
func (s *ThreadSafeSet[T]) Filter(predicate func(T) bool) *ThreadSafeSet[T] {
	s.l.Lock()
	defer s.l.Unlock()
	// we don't lock s2 because it is created here
	s2 := NewThreadSafeSet[T]()
	for k := range s.m {
		if predicate(k) {
			s2.m[k] = struct{}{}
		}
	}
	return s2
}

// Map returns a new set containing the results of applying the function to each element
func (s *ThreadSafeSet[T]) Map(fn func(T) T) *ThreadSafeSet[T] {
	s.l.Lock()
	defer s.l.Unlock()
	// we don't lock s2 because it is created here
	s2 := NewThreadSafeSet[T]()
	for k := range s.m {
		s2.m[fn(k)] = struct{}{}
	}
	return s2
}

// Reduce returns the result of applying the function to each element
func (s *ThreadSafeSet[T]) Reduce(fn func(T, T) T) T {
	s.l.Lock()
	defer s.l.Unlock()
	var result T
	for k := range s.m {
		result = fn(result, k)
	}
	return result
}

// Any returns true if any of the elements in the set pass the predicate
func (s *ThreadSafeSet[T]) Any(predicate func(T) bool) bool {
	s.l.Lock()
	defer s.l.Unlock()
	for k := range s.m {
		if predicate(k) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the set pass the predicate
func (s *ThreadSafeSet[T]) All(predicate func(T) bool) bool {
	s.l.Lock()
	defer s.l.Unlock()
	for k := range s.m {
		if !predicate(k) {
			return false
		}
	}
	return true
}

// String returns a string representation of the set
func (s *ThreadSafeSet[T]) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}
