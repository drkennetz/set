package set

import (
	"sync"
	"testing"
)

func TestNewThreadSafeSet(t *testing.T) {
	s := NewThreadSafeSet[int]()
	if s == nil {
		t.Error("NewThreadSafeSet() returned nil")
	}
}

func TestAdd(t *testing.T) {
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
