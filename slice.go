package main

import (
	"context"
	"errors"
	"reflect"
	"sync"
)

type sliceRepo[T IDer] struct {
	data  []element[T]
	index int
	mu    sync.RWMutex
}

type element[T IDer] struct {
	elem T
	ok   bool
}

func newSliceRepo[T IDer]() *sliceRepo[T] {
	return &sliceRepo[T]{
		data:  make([]element[T], 0),
		index: 0,
		mu:    sync.RWMutex{},
	}
}

func (r *sliceRepo[T]) GetAll(ctx context.Context) []T {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]T, 0, r.index)
	for _, item := range r.data {
		if item.ok {
			result = append(result, item.elem)
		}
	}
	return result
}

func (r *sliceRepo[T]) GetByID(ctx context.Context, id int) (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if id >= len(r.data) {
		var res T
		return res, errors.New("not found")
	}
	res := r.data[id]
	if !res.ok {
		return res.elem, errors.New("not found")
	}
	return res.elem, nil
}

func (r *sliceRepo[T]) Create(ctx context.Context, model T) T {
	r.mu.Lock()
	defer r.mu.Unlock()
	model.SetID(r.index)
	r.data = append(r.data, element[T]{elem: model, ok: true})
	r.index++
	return model
}

func (r *sliceRepo[T]) Remove(ctx context.Context, id int) (T, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if id >= len(r.data) {
		var res T
		return res, errors.New("not found")
	}
	res := r.data[id]
	if !res.ok {
		return res.elem, errors.New("not found")
	}
	r.data[id].ok = false
	return res.elem, nil
}

func (r *sliceRepo[T]) IsEqual(other Repo[T]) bool {
	// TODO: use something better than DeepEqual, which can returns false
	// even if the two Repo are the same (different order).
	return reflect.DeepEqual(r, other)
}
