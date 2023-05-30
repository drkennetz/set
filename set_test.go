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

func TestNewSetFromSlice(t *testing.T) {
	a := NewSetFromSlice[int]([]int{1, 2, 3, 4, 5})
	if !a.Contains(1) {
		t.Error("NewSetFromSlice() failed to add element")
	}
	if !a.Contains(2) {
		t.Error("NewSetFromSlice() failed to add element")
	}
	if !a.Contains(3) {
		t.Error("NewSetFromSlice() failed to add element")
	}
	if !a.Contains(4) {
		t.Error("NewSetFromSlice() failed to add element")
	}
	if !a.Contains(5) {
		t.Error("NewSetFromSlice() failed to add element")
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

func TestSet_Contains(t *testing.T) {
	a := NewSet[string]()
	a.Add("a")
	if !a.Contains("a") {
		t.Error("Set.Contains() failed to determine contains")
	}
	if a.Contains("b") {
		t.Error("Set.Contains() failed to determine contains")
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

func TestSet_Pop(t *testing.T) {
	a := NewSet[int]()
	b := a.Pop()
	if b != 0 {
		t.Error("Set.Pop() failed to return zero value for empty set")
	}
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	a.Add(5)
	c := a.Pop()
	if a.Contains(c) {
		t.Error("Set.Pop() failed to pop element")
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
	c := NewSet[int]()
	c.Add(1)
	c.Add(5)
	if c.IsDisjoint(a) {
		t.Error("Set.IsDisjoint() failed to determine disjoint")
	}
	if c.IsDisjoint(b) {
		t.Error("Set.IsDisjoint() failed to determine disjoint")
	}
}

func TestSet_IsSuperset(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	b := NewSet[int]()
	b.Add(2)
	b.Add(3)
	if !a.IsSuperset(b) {
		t.Error("Set.IsSuperset() failed to determine superset")
	}
	if b.IsSuperset(a) {
		t.Error("Set.IsSuperset() failed to determine superset")
	}
}

func TestSet_IsEqual(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	b := NewSet[int]()
	b.Add(2)
	b.Add(3)
	if a.IsEqual(b) {
		t.Error("Set.IsEqual() failed to determine equality")
	}
	if b.IsEqual(a) {
		t.Error("Set.IsEqual() failed to determine equality")
	}
	c := NewSet[int]()
	c.Add(1)
	c.Add(2)
	c.Add(3)
	c.Add(4)
	if !c.IsEqual(a) {
		t.Error("Set.IsEqual() failed to determine equality")
	}
	if !a.IsEqual(c) {
		t.Error("Set.IsEqual() failed to determine equality")
	}
}

func TestSet_Copy(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	b := a.Copy()
	if !a.IsEqual(b) {
		t.Error("Set.Copy() failed to copy")
	}
	if !b.IsEqual(a) {
		t.Error("Set.Copy() failed to copy")
	}
}

func TestSet_Len(t *testing.T) {
	a := NewSet[string]()
	a.Add("a")
	a.Add("b")
	if a.Len() != 2 {
		t.Error("Set.Len() failed to return correct length")
	}
}

func TestSet_Clear(t *testing.T) {
	a := NewSet[string]()
	a.Add("a")
	a.Add("b")
	a.Clear()
	if a.Len() != 0 {
		t.Error("Set.Clear() failed to clear")
	}
}

func TestSet_IsEmpty(t *testing.T) {
	a := NewSet[string]()
	if !a.IsEmpty() {
		t.Error("Set.IsEmpty() failed to determine empty")
	}
	a.Add("a")
	if a.IsEmpty() {
		t.Error("Set.IsEmpty() failed to determine empty")
	}
}

func TestSet_ToSlice(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	b := a.ToSlice()
	if len(b) != 3 {
		t.Error("Set.ToSlice() failed to convert to slice")
	}
}

func TestSet_Filter(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	b := a.Filter(func(x int) bool {
		return x%2 == 0
	})
	if b.Len() != 2 {
		t.Error("Set.Filter() failed to filter")
	}
	if !b.Contains(2) {
		t.Error("Set.Filter() failed to filter")
	}
	if !b.Contains(4) {
		t.Error("Set.Filter() failed to filter")
	}
	if b.Contains(1) {
		t.Error("Set.Filter() failed to filter")
	}
	if b.Contains(3) {
		t.Error("Set.Filter() failed to filter")
	}
}

func TestSet_Map(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	b := a.Map(func(x int) int {
		return x * 2
	})
	if b.Len() != 4 {
		t.Error("Set.Map() failed to map")
	}
	if !b.Contains(2) {
		t.Error("Set.Map() failed to map")
	}
	if !b.Contains(4) {
		t.Error("Set.Map() failed to map")
	}
	if !b.Contains(6) {
		t.Error("Set.Map() failed to map")
	}
	if !b.Contains(8) {
		t.Error("Set.Map() failed to map")
	}
}

func TestSet_Reduce(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	b := a.Reduce(func(x, y int) int {
		return x + y
	})
	if b != 10 {
		t.Error("Set.Reduce() failed to reduce")
	}
}

func TestSet_Any(t *testing.T) {
	a := NewSet[string]()
	a.Add("cat")
	a.Add("dog")
	a.Add("mouse")
	z := func(x string) bool {
		return x == "dog"
	}
	y := func(x string) bool {
		return x == "bird"
	}
	if !a.Any(z) {
		t.Error("Set.Any() failed to determine any")
	}
	if a.Any(y) {
		t.Error("Set.Any() failed to determine any")
	}
}

func TestSet_All(t *testing.T) {
	a := NewSet[string]()
	a.Add("cat")
	a.Add("dog")
	z := func(x string) bool {
		return len(x) >= 3
	}
	y := func(x string) bool {
		return strings.Contains(x, "a")
	}
	if !a.All(z) {
		t.Error("Set.All() failed to determine all")
	}
	if a.All(y) {
		t.Error("Set.All() failed to determine all")
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
