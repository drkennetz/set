# set
defines the set data structure in go

A set can be instantiated from any [comparable](https://golang.org/ref/spec#Comparison_operators) type.

## Usage

The current implementation of `set` is not concurrency safe, although this addition is welcome / will be added as I continue to develop.
```go
    mySet := set.NewSet[int]()
    mySet.Add(1)
    mySet.Add(2)
    mySet.Remove(2)
    mySet.Contains(1) // true
    mySet.Contains(2) // false
    mySet.Get(2)
... more incoming
```
