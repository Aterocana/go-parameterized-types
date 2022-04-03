package main

import (
	"context"
	"errors"
	"reflect"
	"sync"
)

// mapRepo is an implementation of Repo interface which stores elements in a map.
type mapRepo[T IDer] struct {
	data  map[int]T
	idCnt int
	mu    sync.RWMutex
}

func newMapRepo[T IDer]() *mapRepo[T] {
	return &mapRepo[T]{
		data:  make(map[int]T),
		idCnt: 0,
		mu:    sync.RWMutex{},
	}
}

func (r *mapRepo[T]) GetAll(ctx context.Context) []T {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]T, 0, r.idCnt)
	for _, elem := range r.data {
		result = append(result, elem)
	}
	return result
}

func (r *mapRepo[T]) GetByID(ctx context.Context, id int) (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.data[id]
	if !ok {
		return u, errors.New("not found")
	}
	return u, nil
}

func (r *mapRepo[T]) Create(ctx context.Context, model T) T {
	r.mu.Lock()
	defer r.mu.Unlock()
	model.SetID(r.idCnt)
	r.data[r.idCnt] = model
	r.idCnt++
	return model
}

func (r *mapRepo[T]) Remove(ctx context.Context, id int) (T, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	model, ok := r.data[id]
	if !ok {
		return model, errors.New("not found")
	}
	delete(r.data, id)
	return model, nil
}

func (r *mapRepo[T]) IsEqual(other Repo[T]) bool {
	// TODO: use something better than DeepEqual, which can returns false
	// even if the two Repo are the same (different order).
	return reflect.DeepEqual(r, other)
}
