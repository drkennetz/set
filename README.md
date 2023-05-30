<h1 align="center">set</h1>
<p align="center">
   <a href='#GoVersion'>
      <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/drkennetz/set">
   </a>
    <a href="https://github.com/drkennetz/set">
        <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/drkennetz/set/ci.yml">
    </a>
    <a href="https://codecov.io/github/drkennetz/set" >
        <img src="https://codecov.io/github/drkennetz/set/branch/main/graph/badge.svg?token=8AVIX80EMG"/>
    </a>
    <a href="https://github.com/drkennetz/set/issues">
        <img src="https://img.shields.io/github/issues/drkennetz/set" alt="Issues">
   </a>
</p>

Provides a set data structure for Go.

A set can be instantiated from any [comparable](https://golang.org/ref/spec#Comparison_operators) type.

## Usage

The current implementation of `set` is not concurrency safe, although this addition is welcome / will be added as I continue to develop.

## Set Basics
```go
    mySet := set.NewSet[int]()
	// alternatively, from a slice
	// mySet := NewSetFromSlice[int]([]int{1, 2})
    mySet.Add(1)
    mySet.Add(2)
	// Remove 2 from set
    mySet.Remove(2)
	// Check if set contains 1
    mySet.Contains(1) // true
    mySet.Contains(2) // false
	// Length of set
    mySet.Len() // 1
	// Clear the set
    mySet.Clear()
	mySet.Len() // 0
	// Pop an arbitrary element from the set
	mySet.Pop() // 0
	mySet.Add(1)
	a := mySet.Pop() // 1, 1 is removed from set
```

## Common Set Operations
```go
    setA := set.NewSet[int]()
    setA.Add(1)
    setA.Add(2)
    setA.Add(3)

    setB := set.NewSet[int]()
    setB.Add(3)
    setB.Add(4)
    setB.Add(5)

    setC := setA.Union(setB) // {1, 2, 3, 4, 5}
    setD := setA.Intersect(setB) // {3}
    setE := setA.Difference(setB) // {1, 2}
    setF := setA.SymmetricDifference(setB) // {1, 2, 4, 5}
```

## Common Set Helpers
```go
    setA := set.NewSet[int]()
    setA.Add(3)

    setB := set.NewSet[int]()
    setB.Add(3)
    setB.Add(4)
    setB.Add(5)

	// Checks if setA is a subset of setB (setA ⊆ setB)
    setA.IsSubset(setB) // true
	// Checks if setA is a superset of setB (setA ⊇ setB)
    setA.IsSuperset(setB) // false
	// Checks if setA is disjoint from setB (setA ∩ setB = ∅)
    setA.IsDisjoint(setB) // false
	// Checks if setA is equal to setB (setA == setB)
    setA.IsEqual(setB) // false
	// Checks if setA is empty (|setA| = 0)
    setA.IsEmpty() // false
	// Converts setA to a slice, arbitrary order
    setA.ToSlice() // []int{3}
	// Copies setA to a new set
	setC := setA.Copy() // {3}
	// String representation of setA
	setA.String() // "{3}"
```

## Slightly More Advanced Set Methods
```go
    setA := set.NewSet[string]()
	setA.Add("cat")
	setA.Add("dog")
	setA.Add("fish")
	lenThree := func(s string) bool {
        return len(s) == 3
    }
	// Filter setA by lenThree
	setB := setA.Filter(lenThree) // {"cat", "dog"}
	// Map setA to a new set
	setC := setA.Map(func(s string) string {
        return s + "s"
    }) // {"cats", "dogs", "fishs"}
	// Reduce setA to a single value
	setD := setA.Reduce(func(s1, s2 string) string {
        return s1 + s2
    }) // "catdogfish"
	// Check if all elements in setA satisfy a predicate
	setA.All(func(s string) bool {
        return len(s) > 1
    }) // true
	// Check if any elements in setA satisfy a predicate
	setA.Any(func(s string) bool {
        return len(s) > 3
    }) // true
```

## Contributing
