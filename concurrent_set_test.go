package set

import (
	"strconv"
	"strings"
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

// go test -run TestNewThreadSafeSetFromSlice .
func TestTSSNewThreadSafeSetFromSlice(t *testing.T) {
	s := NewThreadSafeSetFromSlice([]int{1, 2, 3, 4, 5})
	if s == nil {
		t.Error("NewThreadSafeSetFromSlice() returned nil")
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
	zero := make(chan int, 1)
	go func() {
		zero <- s.Pop()
	}()
	if <-zero != 0 {
		t.Errorf("Expected 0, got %d", <-zero)
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

// go test -race -run TestTSSSymmetricDifference .
func TestTSSSymmetricDifference(t *testing.T) {
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
		v := s.SymmetricDifference(s2)
		ch <- v
	}()
	go func() {
		v := s2.SymmetricDifference(s)
		ch <- v
	}()
	// test for correctness
	for i := 0; i < 2; i++ {
		v := <-ch
		for j := 20; j < 30; j++ {
			if !v.Contains(j) {
				t.Errorf("Expected true, got false")
			}
		}
	}
}

// go test -race -run TestTSSIsSubset .
func TestTSSIsSubset(t *testing.T) {
	s := NewThreadSafeSet[int]()
	s2 := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	for i := 10; i < 30; i++ {
		s2.Add(i)
	}
	ch := make(chan bool, 1)
	go func() {
		v := s.IsSubset(s2)
		ch <- v
	}()
	// test for correctness
	v := <-ch
	if !v {
		t.Errorf("Expected true, got false")
	}
	go func() {
		v := s2.IsSubset(s)
		ch <- v
	}()
	// test for correctness
	v = <-ch
	if v {
		t.Errorf("Expected false, got true")
	}
}

// go test -race -run TestTSSIsSuperset .
func TestTSSIsSuperset(t *testing.T) {
	s := NewThreadSafeSet[int]()
	s2 := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	for i := 10; i < 30; i++ {
		s2.Add(i)
	}
	ch := make(chan bool, 1)
	go func() {
		v := s2.IsSuperset(s)
		ch <- v
	}()
	// test for correctness
	v := <-ch
	if !v {
		t.Errorf("Expected true, got false")
	}
}

// go test -race -run TestTSSIsDisjoint .
func TestTSSIsDisjoint(t *testing.T) {
	s := NewThreadSafeSet[int]()
	s2 := NewThreadSafeSet[int]()
	s3 := NewThreadSafeSet[int]()
	for i := 20; i < 40; i++ {
		s.Add(i)
	}
	for i := 10; i < 20; i++ {
		s2.Add(i)
	}
	for i := 30; i < 40; i++ {
		s3.Add(i)
	}
	ch := make(chan bool, 1)
	go func() {
		v := s.IsDisjoint(s2)
		ch <- v
	}()
	// test for correctness
	v := <-ch
	if !v {
		t.Errorf("Expected true, got false")
	}
	go func() {
		v := s.IsDisjoint(s3)
		ch <- v
	}()
	// test for correctness
	v = <-ch
	if v {
		t.Errorf("Expected false, got true")
	}
}

// go test -race -run TestTSSIsEqual .
func TestTSSIsEqual(t *testing.T) {
	s := NewThreadSafeSet[int]()
	s2 := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	for i := 10; i < 20; i++ {
		s2.Add(i)
	}
	ch := make(chan bool, 1)
	go func() {
		v := s.IsEqual(s2)
		ch <- v
	}()
	// test for correctness
	v := <-ch
	if !v {
		t.Errorf("Expected true, got false")
	}
}

// go test -race -run TestTSSCopy .
func TestTSSCopy(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	ch := make(chan *ThreadSafeSet[int], 1)
	go func() {
		v := s.Copy()
		ch <- v
	}()
	// test for correctness
	v := <-ch
	if !v.IsEqual(s) {
		t.Errorf("Expected true, got false")
	}
}

// go test -race -run TestTSSLen .
func TestTSSLen(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	ch := make(chan int, 1)
	go func() {
		v := s.Len()
		ch <- v
	}()
	// test for correctness
	v := <-ch
	if v != 10 {
		t.Errorf("Expected true, got false")
	}
}

// go test -race -run TestTSSClear .
func TestTSSClear(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	ch := make(chan bool, 1)
	go func() {
		s.Clear()
		ch <- true
	}()
	// test for correctness
	<-ch
	if s.Len() != 0 {
		t.Errorf("Expected true, got false")
	}
}

// go test -race -run TestTSSIsEmpty .
func TestTSSIsEmpty(t *testing.T) {
	s := NewThreadSafeSet[int]()
	ch := make(chan bool, 1)
	go func() {
		v := s.IsEmpty()
		ch <- v
	}()
	// test for correctness
	v := <-ch
	if !v {
		t.Errorf("Expected true, got false")
	}
}

// go test -race -run TestTSSToSlice .
func TestTSSToSlice(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	ch := make(chan []int, 1)
	go func() {
		v := s.ToSlice()
		ch <- v
	}()
	// test for correctness (kind of)
	v := <-ch
	if len(v) != 10 {
		t.Errorf("Expected true, got false")
	}
}

// go test -race -run TestTSSFilter .
func TestTSSFilter(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	ch := make(chan *ThreadSafeSet[int], 1)
	go func() {
		v := s.Filter(func(i int) bool {
			return i%2 == 0
		})
		ch <- v
	}()
	// test for correctness
	v := <-ch
	for key := range v.m {
		if key%2 != 0 {
			t.Errorf("Expected true, got false")
		}
	}
}

// go test -race -run TestTSSMap .
func TestTSSMap(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 10; i < 20; i++ {
		s.Add(i)
	}
	ch := make(chan *ThreadSafeSet[int], 1)
	go func() {
		v := s.Map(func(i int) int {
			return i * 2
		})
		ch <- v
	}()
	// test for correctness
	v := <-ch
	for key := range v.m {
		if key%2 != 0 {
			t.Errorf("Expected true, got false")
		}
	}
}

// go test -race -run TestTSSReduce .
func TestTSSReduce(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 1; i < 10; i++ {
		s.Add(i)
	}
	ch := make(chan int, 1)
	go func() {
		v := s.Reduce(func(acc, i int) int {
			return acc + i
		})
		ch <- v
	}()
	// test for correctness
	v := <-ch
	if v != 45 {
		t.Errorf("Expected true, got false")
	}
}

// go test -race -run TestTSSAny .
func TestTSSAny(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 1; i < 10; i++ {
		s.Add(i)
	}
	ch := make(chan bool, 1)
	go func() {
		v := s.Any(func(i int) bool {
			return i == 5
		})
		ch <- v
	}()
	// test for correctness
	v := <-ch
	if !v {
		t.Errorf("Expected true, got false")
	}
	go func() {
		v := s.Any(func(i int) bool {
			return i == 10
		})
		ch <- v
	}()
	// test for correctness
	v = <-ch
	if v {
		t.Errorf("Expected false, got true")
	}
}

// go test -race -run TestTSSAll .
func TestTSSAll(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 1; i < 10; i++ {
		s.Add(i)
	}
	ch := make(chan bool, 1)
	go func() {
		v := s.All(func(i int) bool {
			return i < 10
		})
		ch <- v
	}()
	// test for correctness
	v := <-ch
	if !v {
		t.Errorf("Expected true, got false")
	}
	go func() {
		v := s.All(func(i int) bool {
			return i < 5
		})
		ch <- v
	}()
	// test for correctness
	v = <-ch
	if v {
		t.Errorf("Expected false, got true")
	}
}

// go test -race -run TestTSSString .
func TestTSSString(t *testing.T) {
	s := NewThreadSafeSet[int]()
	for i := 1; i < 10; i++ {
		s.Add(i)
	}
	ch := make(chan string, 1)
	go func() {
		v := s.String()
		ch <- v
	}()
	// test for correctness
	v := <-ch

	// check for ints in any order because order is not guaranteed
	for i := 1; i < 10; i++ {
		if !strings.Contains(v, strconv.Itoa(i)) {
			t.Errorf("Expected true, got false")
		}
	}
}
