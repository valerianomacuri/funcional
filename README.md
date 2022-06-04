# funcional

[![Go Reference](https://pkg.go.dev/badge/github.com/valerianomacuri/funcional.svg)](https://pkg.go.dev/github.com/valerianomacuri/funcional)

A very useful library to manipulate data structures in Go, let's try it Gophers.
funcional uses generics, released in Go 1.18, for easy reading.

## Check the testing

```console
go test -coverprofile cover.out
```

```console
go tool cover -func cover.out
```

## Install

```console
go get -u github.com/valerianomacuri/funcional
```

## Examples

### Filter

```go
filteredValues := Filter([]int{1, 2, 3, 4}, func(num int) bool {
		if num == 1 {
			return false
		} else {
			return true
		}
	})
```

### Find

```go
slice := []int{1, 2, 3, 4}
foundValue, ok := Find(slice, func(value int) bool {
	return value == 5
})
```

### Includes

```go
slice := []int{1, 3, 4, 5}
isIncluded := Includes(slice, 1)
```

### Map

```go
result := Map([]int{1, 2, 3, 4, 5}, func(value int) int {
            return value + 1
        })
```

### Reduce

```go
slice := []int{1, 2, 3, 4, 5}
reducedValue := Reduce(slice, 0, func(num1, num2 int) int {
    return num1 + num2
})
```

### Some

```go
slice := []int{1, 3, 4, 5}
result := Some(slice, func(value int) bool {
    return (value % 2) == 0
})
```
