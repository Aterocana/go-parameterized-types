package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

// utils_test.go contains some utilities functions and data structures for
// uniform tests across all Repo[T] implementations.

func testName(name, structName string) string {
	return fmt.Sprintf("%s [%s]", name, structName)
}

func areRepoEquals[T IDer](r1, r2 Repo[T]) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	return r1.IsEqual(r2)
}

type getAllTest[T IDer] struct {
	name     string
	repo     Repo[T]
	expected []T
}

func testGetAll[T IDer](structName string, t *testing.T, tests []getAllTest[T]) {
	for _, tt := range tests {
		ctx := context.TODO()
		t.Run(testName(tt.name, structName), func(t *testing.T) {
			got := tt.repo.GetAll(ctx)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("%s.GetAll() = expected %v, got %v", structName, tt.expected, got)
				return
			}
		})
	}
}

type getByIDTest[T IDer] struct {
	name string
	repo Repo[T]
	id   int
	err  error
}

func testGetByID[T IDer](structName string, t *testing.T, tests []getByIDTest[T]) {
	for _, tt := range tests {
		ctx := context.TODO()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetByID(ctx, tt.id)
			if err != nil {
				if tt.err == nil {
					t.Errorf("%s.GetByID() = unexpected error %v", structName, err)
					return
				}
				if tt.err.Error() != err.Error() {
					t.Errorf("%s.GetByID() = expected error %v, got %v", structName, tt.err, err)
					return
				}
				return
			}
			if tt.err != nil {
				t.Errorf("%s.GetByID() = expected error %v, got none", structName, tt.err)
				return
			}
			if got.GetID() != tt.id {
				t.Errorf("%s.GetByID() = expected id %v, got %v", structName, tt.id, got.GetID())
				return
			}
		})
	}
}

type createTest[T IDer] struct {
	name     string
	repo     Repo[T]
	item     T
	expected Repo[T]
	id       int
}

func testCreate[T IDer](structName string, t *testing.T, tests []createTest[T]) {
	for _, tt := range tests {
		ctx := context.TODO()
		t.Run(testName(tt.name, structName), func(t *testing.T) {
			got := tt.repo.Create(ctx, tt.item)
			if got.GetID() != tt.id {
				t.Errorf("%s.Create() = expected id %v, got %v", structName, tt.id, got.GetID())
				return
			}
			if !areRepoEquals(tt.repo, tt.expected) {
				t.Errorf("%s.Create() = expected %v, got %v", structName, tt.expected, tt.repo)
				return
			}
		})
	}
}

type removeTest[T IDer] struct {
	name        string
	repo        Repo[T]
	id          int
	expected    Repo[T]
	removedItem T
	err         error
}

func testRemove[T IDer](structName string, t *testing.T, tests []removeTest[T]) {
	for _, tt := range tests {
		ctx := context.TODO()
		t.Run(testName(tt.name, structName), func(t *testing.T) {
			got, err := tt.repo.Remove(ctx, tt.id)
			if err != nil {
				if tt.err == nil {
					t.Errorf("%s.Remove() = unexpected error %v", structName, err)
					return
				}
				if tt.err.Error() != err.Error() {
					t.Errorf("%s.Remove() = expected error %v, got %v", structName, tt.err, err)
					return
				}
				return
			}
			if tt.err != nil {
				t.Errorf("%s.Remove() = expected error %v, got none", structName, tt.err)
				return
			}
			if !reflect.DeepEqual(tt.removedItem, got) {
				t.Errorf("%s.Remove() = expected %v, got %v", structName, tt.removedItem, got)
				return
			}
			if !areRepoEquals(tt.repo, tt.expected) {
				t.Errorf("%s.Remove() = expected %v, got %v", structName, tt.expected, tt.repo)
				return
			}
		})
	}
}
