package set

import "testing"

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
