package main

import (
	"context"
	"log"
)

// IDer is a type containing an ID which can be read and written.
type IDer interface {
	GetID() int
	SetID(id int)
}

// Repo is something storing T objects
type Repo[T IDer] interface {
	// GetAll should returns all Repo elements
	GetAll(ctx context.Context) []T

	// GetByID should returns an element in the Repo by its ID, if possible.
	// It should return an error if the element is not found.
	GetByID(ctx context.Context, id int) (T, error)

	// Create should create a new element in the Repo and returns it
	Create(ctx context.Context, model T) T

	// Remove should remove an element from the Repo by its ID, if possible.
	// It should return an error if the element is not found.
	Remove(ctx context.Context, id int) (T, error)

	// IsEqual should returns true if current Repo is equal to the provided one.
	// It is used for testing purposes.
	IsEqual(other Repo[T]) bool
}

// PrintAll logs alla elements of a provided Repo
func PrintAll[T IDer](repo Repo[T]) {
	log.Println(repo.GetAll(context.Background()))
}
