package main

import (
	"errors"
	"testing"
)

func TestMapGetAll(t *testing.T) {
	userTests := []getAllTest[*user]{
		{
			name: "empty map",
			repo: &mapRepo[*user]{
				idCnt: 0,
				data:  map[int]*user{},
			},
			expected: []*user{},
		},
		{
			name: "some elements",
			repo: &mapRepo[*user]{
				idCnt: 2,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
					1: {id: 1, name: "user1"},
				},
			},
			expected: []*user{
				{id: 0, name: "user0"},
				{id: 1, name: "user1"},
			},
		},
	}
	dogTests := []getAllTest[*user]{
		{
			name: "empty map",
			repo: &mapRepo[*user]{
				idCnt: 0,
				data:  map[int]*user{},
			},
			expected: []*user{},
		},
		{
			name: "some elements",
			repo: &mapRepo[*user]{
				idCnt: 2,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
					1: {id: 1, name: "user1"},
				},
			},
			expected: []*user{
				{id: 0, name: "user0"},
				{id: 1, name: "user1"},
			},
		},
	}
	testGetAll("mapRepo[*user]", t, userTests)
	testGetAll("mapRepo[*dog]", t, dogTests)
}

func TestMapGetByID(t *testing.T) {
	userTests := []getByIDTest[*user]{
		{
			name: "user map, id present",
			repo: &mapRepo[*user]{
				idCnt: 1,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
				},
			},
			id: 0,
		},
		{
			name: "user map, id not present",
			repo: &mapRepo[*user]{
				idCnt: 1,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
				},
			},
			id:  1,
			err: errors.New("not found"),
		},
		{
			name: "empty user map",
			repo: &mapRepo[*user]{
				idCnt: 0,
				data:  map[int]*user{},
			},
			id:  0,
			err: errors.New("not found"),
		},
	}
	dogTests := []getByIDTest[*dog]{
		{
			name: "dog map, id present",
			repo: &mapRepo[*dog]{
				idCnt: 1,
				data: map[int]*dog{
					0: {id: 0, name: "dog0"},
				},
			},
			id: 0,
		},
		{
			name: "dog map, id not present",
			repo: &mapRepo[*dog]{
				idCnt: 1,
				data: map[int]*dog{
					0: {id: 0, name: "dog0"},
				},
			},
			id:  1,
			err: errors.New("not found"),
		},
		{
			name: "empty dog map",
			repo: &mapRepo[*dog]{
				idCnt: 0,
				data:  map[int]*dog{},
			},
			id:  0,
			err: errors.New("not found"),
		},
	}
	testGetByID("mapRepo[*user]", t, userTests)
	testGetByID("mapRepo[*dog]", t, dogTests)
}

func TestMapCreate(t *testing.T) {
	userTests := []createTest[*user]{
		{
			name: "empty map",
			repo: &mapRepo[*user]{
				idCnt: 0,
				data:  map[int]*user{},
			},
			expected: &mapRepo[*user]{
				idCnt: 1,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
				},
			},
			item: &user{name: "user0"},
			id:   0,
		},
		{
			name: "some elements map",
			repo: &mapRepo[*user]{
				idCnt: 2,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
					1: {id: 1, name: "user1"},
				},
			},
			expected: &mapRepo[*user]{
				idCnt: 3,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
					1: {id: 1, name: "user1"},
					2: {id: 2, name: "user2"},
				},
			},
			item: &user{name: "user2"},
			id:   2,
		},
	}
	dogTests := []createTest[*user]{
		{
			name: "empty map",
			repo: &mapRepo[*user]{
				idCnt: 0,
				data:  map[int]*user{},
			},
			expected: &mapRepo[*user]{
				idCnt: 1,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
				},
			},
			item: &user{name: "user0"},
			id:   0,
		},
		{
			name: "some elements map",
			repo: &mapRepo[*user]{
				idCnt: 2,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
					1: {id: 1, name: "user1"},
				},
			},
			expected: &mapRepo[*user]{
				idCnt: 3,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
					1: {id: 1, name: "user1"},
					2: {id: 2, name: "user2"},
				},
			},
			item: &user{name: "user2"},
			id:   2,
		},
	}
	testCreate("mapRepo[*user]", t, userTests)
	testCreate("mapRepo[*dog]", t, dogTests)
}

func TestMapRemove(t *testing.T) {
	userTests := []removeTest[*user]{
		{
			name: "empty map",
			repo: &mapRepo[*user]{
				idCnt: 0,
				data:  map[int]*user{},
			},
			id:  0,
			err: errors.New("not found"),
		},
		{
			name: "id out of range",
			repo: &mapRepo[*user]{
				idCnt: 2,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
					1: {id: 1, name: "user1"},
				},
			},
			id:  2,
			err: errors.New("not found"),
		},
		{
			name: "element removed",
			repo: &mapRepo[*user]{
				idCnt: 2,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
					1: {id: 1, name: "user1"},
				},
			},
			id:          1,
			removedItem: &user{id: 1, name: "user1"},
			expected: &mapRepo[*user]{
				idCnt: 2,
				data: map[int]*user{
					0: {id: 0, name: "user0"},
				},
			},
		},
	}
	dogTests := []removeTest[*dog]{
		{
			name: "empty map",
			repo: &mapRepo[*dog]{
				idCnt: 0,
				data:  map[int]*dog{},
			},
			id:  0,
			err: errors.New("not found"),
		},
		{
			name: "id out of range",
			repo: &mapRepo[*dog]{
				idCnt: 2,
				data: map[int]*dog{
					0: {id: 0, name: "dog0"},
					1: {id: 1, name: "dog1"},
				},
			},
			id:  2,
			err: errors.New("not found"),
		},
		{
			name: "element removed",
			repo: &mapRepo[*dog]{
				idCnt: 2,
				data: map[int]*dog{
					0: {id: 0, name: "dog0"},
					1: {id: 1, name: "dog1"},
				},
			},
			id:          1,
			removedItem: &dog{id: 1, name: "dog1"},
			expected: &mapRepo[*dog]{
				idCnt: 2,
				data: map[int]*dog{
					0: {id: 0, name: "dog0"},
				},
			},
		},
	}
	testRemove("mapRepo[*user]", t, userTests)
	testRemove("mapRepo[*dog]", t, dogTests)
}
