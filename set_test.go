package set

import (
	"strings"
	"testing"
)

func TestNewSet(t *testing.T) {
	a := NewSet[int]()
	if a == nil {
		t.Error("NewSet() returned nil")
	}
	b := NewSet[string]()
	if b == nil {
		t.Error("NewSet() returned nil")
	}
	c := NewSet[float64]()
	if c == nil {
		t.Error("NewSet() returned nil")
	}
	d := NewSet[bool]()
	if d == nil {
		t.Error("NewSet() returned nil")
	}
	e := NewSet[complex128]()
	if e == nil {
		t.Error("NewSet() returned nil")
	}
	f := NewSet[complex64]()
	if f == nil {
		t.Error("NewSet() returned nil")
	}
	g := NewSet[uint]()
	if g == nil {
		t.Error("NewSet() returned nil")
	}
	h := NewSet[uint8]()
	if h == nil {
		t.Error("NewSet() returned nil")
	}
	i := NewSet[uint16]()
	if i == nil {
		t.Error("NewSet() returned nil")
	}
	j := NewSet[uint32]()
	if j == nil {
		t.Error("NewSet() returned nil")
	}
	k := NewSet[uint64]()
	if k == nil {
		t.Error("NewSet() returned nil")
	}
	l := NewSet[int8]()
	if l == nil {
		t.Error("NewSet() returned nil")
	}
	m := NewSet[int16]()
	if m == nil {
		t.Error("NewSet() returned nil")
	}
	n := NewSet[int32]()
	if n == nil {
		t.Error("NewSet() returned nil")
	}
	o := NewSet[int64]()
	if o == nil {
		t.Error("NewSet() returned nil")
	}
	p := NewSet[uintptr]()
	if p == nil {
		t.Error("NewSet() returned nil")
	}
	q := NewSet[rune]()
	if q == nil {
		t.Error("NewSet() returned nil")
	}
	r := NewSet[byte]()
	if r == nil {
		t.Error("NewSet() returned nil")
	}
	s := NewSet[error]()
	if s == nil {
		t.Error("NewSet() returned nil")
	}
	tt := NewSet[interface{}]()
	if tt == nil {
		t.Error("NewSet() returned nil")
	}
	u := NewSet[struct{}]()
	if u == nil {
		t.Error("NewSet() returned nil")
	}
	v := NewSet[[1]int]()
	if v == nil {
		t.Error("NewSet() returned nil")
	}
	w := NewSet[*int]()
	if w == nil {
		t.Error("NewSet() returned nil")
	}
	x := NewSet[chan int]()
	if x == nil {
		t.Error("NewSet() returned nil")
	}
	// testing a custom type with underlying type comparable
	type hotdog int
	y := NewSet[hotdog]()
	if y == nil {
		t.Error("NewSet() returned nil")
	}
}

func TestSet_Add(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	if !a.Contains(1) {
		t.Error("Set.Add() failed to add element")
	}
	a.Add(2)
	if !a.Contains(2) {
		t.Error("Set.Add() failed to add element")
	}
	a.Add(3)
	if !a.Contains(3) {
		t.Error("Set.Add() failed to add element")
	}
	a.Add(4)
	if !a.Contains(4) {
		t.Error("Set.Add() failed to add element")
	}
	a.Add(5)
	if !a.Contains(5) {
		t.Error("Set.Add() failed to add element")
	}
}

func TestSet_Remove(t *testing.T) {
	a := NewSet[string]()
	a.Add("a")
	a.Remove("a")
	if a.Contains("a") {
		t.Error("Set.Remove() failed to remove element")
	}
}

func TestSet_Intersection(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	a.Add(5)
	b := NewSet[int]()
	b.Add(2)
	b.Add(3)
	c := a.Intersection(b)
	if !c.Contains(2) {
		t.Error("Set.Intersection() failed to intersect")
	}
	if !c.Contains(3) {
		t.Error("Set.Intersection() failed to intersect")
	}
	if c.Contains(4) {
		t.Error("Set.Intersection() failed to intersect")
	}
	if c.Contains(1) {
		t.Error("Set.Intersection() failed to intersect")
	}
	if c.Contains(5) {
		t.Error("Set.Intersection() failed to intersect")
	}
}

func TestSet_Union(t *testing.T) {
	a := NewSet[string]()
	a.Add("a")
	a.Add("b")
	a.Add("c")
	b := NewSet[string]()
	b.Add("b")
	b.Add("c")
	b.Add("d")
	c := a.Union(b)
	if !c.Contains("a") {
		t.Error("Set.Union() failed to union")
	}
	if !c.Contains("b") {
		t.Error("Set.Union() failed to union")
	}
	if !c.Contains("c") {
		t.Error("Set.Union() failed to union")
	}
	if !c.Contains("d") {
		t.Error("Set.Union() failed to union")
	}
}

func TestSet_Difference(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	b := NewSet[int]()
	b.Add(2)
	b.Add(3)
	b.Add(5)
	c := a.Difference(b)
	if c.Contains(2) {
		t.Error("Set.Difference() failed to difference")
	}
	if c.Contains(3) {
		t.Error("Set.Difference() failed to difference")
	}
	if !c.Contains(4) {
		t.Error("Set.Difference() failed to difference")
	}
	if !c.Contains(1) {
		t.Error("Set.Difference() failed to difference")
	}
	if c.Contains(5) {
		t.Error("Set.Difference() failed to difference")
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	b := NewSet[int]()
	b.Add(2)
	b.Add(3)
	b.Add(5)
	c := a.SymmetricDifference(b)
	if c.Contains(2) {
		t.Error("Set.SymmetricDifference() failed to symmetric difference")
	}
	if c.Contains(3) {
		t.Error("Set.SymmetricDifference() failed to symmetric difference")
	}
	if !c.Contains(4) {
		t.Error("Set.SymmetricDifference() failed to symmetric difference")
	}
	if !c.Contains(1) {
		t.Error("Set.SymmetricDifference() failed to symmetric difference")
	}
	if !c.Contains(5) {
		t.Error("Set.SymmetricDifference() failed to symmetric difference")
	}
}

func TestSet_IsSubset(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	b := NewSet[int]()
	b.Add(2)
	b.Add(3)
	if !b.IsSubset(a) {
		t.Error("Set.IsSubset() failed to determine subset")
	}
	if a.IsSubset(b) {
		t.Error("Set.IsSubset() failed to determine subset")
	}
}

func TestSet_IsDisjoint(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	b := NewSet[int]()
	b.Add(5)
	b.Add(6)
	if !b.IsDisjoint(a) {
		t.Error("Set.IsDisjoint() failed to determine disjoint")
	}
	if !a.IsDisjoint(b) {
		t.Error("Set.IsDisjoint() failed to determine disjoint")
	}
}

func TestSet_String(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	// order is not guaranteed
	s := a.String()
	if !strings.Contains(s, "1") {
		t.Error("Set.String() failed to stringify")
	}
	if !strings.Contains(s, "2") {
		t.Error("Set.String() failed to stringify")
	}
	if !strings.Contains(s, "3") {
		t.Error("Set.String() failed to stringify")
	}
}
