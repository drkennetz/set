package set

import (
	"sync"
	"testing"
)

// go test -run TestNewThreadSafeSet .
func TestTSSNewThreadSafeSet(t *testing.T) {
	s := NewThreadSafeSet[int]()
	if s == nil {
		t.Error("NewThreadSafeSet() returned nil")
	}
}

// go test -race -run TestTSSAdd .
func TestTSSAdd(t *testing.T) {
	s := NewThreadSafeSet[int]()
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(i)
		go func(i int) {
			defer wg.Done()
			s.Add(i + 10)
		}(i)
	}
	wg.Wait()
	if s.Len() != 20 {
		t.Errorf("Expected length 20, got %d", s.Len())
	}
}

// go test -race -run TestTSSContains .
func TestTSSContains(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 0; i < 20; i++ {
		s.Add(i)
	}
	bools := make(chan bool, 3)
	go func() {
		bools <- s.Contains(10)
	}()
	go func() {
		bools <- s.Contains(2)
	}()
	go func() {
		bools <- s.Contains(0)
	}()
	for i := 0; i < 3; i++ {
		if !<-bools {
			t.Errorf("Expected true, got false")
		}
	}
}

// go test -race -run TestTSSRemove .
func TestTSSRemove(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 0; i < 20; i++ {
		s.Add(i)
	}
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func(i int) {
			defer wg.Done()
			s.Remove(i)
		}(i)
	}
	wg.Wait()
	if !s.IsEmpty() {
		t.Errorf("Expected empty set, got %v", s)
	}
}

// go test -race -run TestTSSPop .
func TestTSSPop(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 0; i < 20; i++ {
		s.Add(i)
	}
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			s.Pop()
		}()
	}
	wg.Wait()
	if !s.IsEmpty() {
		t.Errorf("Expected empty set, got %v", s)
	}
}

// go test -race -run TestTSSIntersection .
func TestTSSIntersection(t *testing.T) {
	s := NewThreadSafeSet[int]()
	s2 := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	for i := 10; i < 30; i++ {
		s2.Add(i)
	}
	ch := make(chan *ThreadSafeSet[int], 2)
	go func() {
		v := s.Intersection(s2)
		ch <- v
	}()
	go func() {
		v := s2.Intersection(s)
		ch <- v
	}()
	// test for correctness
	for i := 0; i < 2; i++ {
		v := <-ch
		for j := 10; j < 20; j++ {
			if !v.Contains(j) {
				t.Errorf("Expected true, got false")
			}
		}
	}
}

// go test -race -run TestTSSUnion .
func TestTSSUnion(t *testing.T) {
	s := NewThreadSafeSet[int]()
	s2 := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	for i := 10; i < 30; i++ {
		s2.Add(i)
	}
	ch := make(chan *ThreadSafeSet[int], 2)
	go func() {
		v := s.Union(s2)
		ch <- v
	}()
	go func() {
		v := s2.Union(s)
		ch <- v
	}()
	// test for correctness
	for i := 0; i < 2; i++ {
		v := <-ch
		for j := 10; j < 30; j++ {
			if !v.Contains(j) {
				t.Errorf("Expected true, got false")
			}
		}
	}
}

// go test -race -run TestTSSDifference .
func TestTSSDifference(t *testing.T) {
	s := NewThreadSafeSet[int]()
	s2 := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	for i := 10; i < 30; i++ {
		s2.Add(i)
	}
	ch := make(chan *ThreadSafeSet[int], 2)
	go func() {
		v := s.Difference(s2)
		ch <- v
	}()
	go func() {
		v := s2.Difference(s)
		ch <- v
	}()
	// test for correctness
	for i := 0; i < 2; i++ {
		v := <-ch
		if v.Len() != 0 {
			// one set is a subset of the other
			// so one should have a difference, and one should not
			for j := 20; j < 30; j++ {
				if !v.Contains(j) {
					t.Errorf("Expected true, got false")
				}
			}
		}
	}
}
