# go-parameterized-types

This repo contains some tests written for getting my feet wet with 1.18 parameterized types new feature, used as a proof of concept.

Use go compiler with minimum version [1.18](https://go.dev/blog/go1.18).

## Index

* [Synopsis](#synopsis)
  * [Repo[T] definition](#repot-definition)
  * [Repo[T] implementations](#repot-implementations)
  * [Tests](#tests)
* [Makefile commands](#makefile)
  * [Build](#build)
  * [Test](#test)
  * [Run](#run)

## Synopsis

### Repo[T] definition

File `repo.go` contains the definition of a generic `Repo` type, which represents something that can store items. `Repo` is parameterized by the type `T` (item type), where `T` is defined as following:
```go
type IDer interface {
	GetID() int
	SetID(id int)
}
```

so basically anything containing an int ID which can be set and retrieved.

### Repo[T] implementations

Two simple implementations are provided:

* a `map[int]T` called `mapRepo` in file `map.go`;
* a `slice` called `sliceRepo` in file `slice.go`.


### Tests

Tests are also written in a "generic syntax", avoiding code duplication with functions and structs implemented in `utils_test.go`. For each `Repo` method, a test struct and a test function are implemented. Then those are used in `map_test.go` and `slice_test.go` to test the implemented methods.

## Makefile

`Makefile` provides all the necessary tools to run the tests and the simple `main` provided as example.
Run `make init-git` to setup local git repo with all githooks in `.tools/cmd/githooks`.

Available commands are:

### Build

`make build` to build the binary;

### Test

`make test` to run tests;

### Run

`make run` to run the binary.
