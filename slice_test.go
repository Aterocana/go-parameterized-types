package main

import (
	"errors"
	"testing"
)

func TestSliceGetAll(t *testing.T) {
	userTests := []getAllTest[*user]{
		{
			name: "empty slice",
			repo: &sliceRepo[*user]{
				index: 0,
				data:  []element[*user]{},
			},
			expected: []*user{},
		},
		{
			name: "all ok elements",
			repo: &sliceRepo[*user]{
				index: 2,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
					{
						elem: &user{id: 1, name: "user1"},
						ok:   true,
					},
				},
			},
			expected: []*user{
				{id: 0, name: "user0"},
				{id: 1, name: "user1"},
			},
		},
		{
			name: "some removed elements",
			repo: &sliceRepo[*user]{
				index: 3,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
					{
						elem: &user{id: 1, name: "user1"},
						ok:   false,
					},
					{
						elem: &user{id: 2, name: "user2"},
						ok:   true,
					},
				},
			},
			expected: []*user{
				{id: 0, name: "user0"},
				{id: 2, name: "user2"},
			},
		},
		{
			name: "all removed elements",
			repo: &sliceRepo[*user]{
				index: 3,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   false,
					},
					{
						elem: &user{id: 1, name: "user1"},
						ok:   false,
					},
					{
						elem: &user{id: 2, name: "user2"},
						ok:   false,
					},
				},
			},
			expected: []*user{},
		},
	}
	dogTests := []getAllTest[*dog]{
		{
			name: "empty slice",
			repo: &sliceRepo[*dog]{
				index: 0,
				data:  []element[*dog]{},
			},
			expected: []*dog{},
		},
		{
			name: "all ok elements",
			repo: &sliceRepo[*dog]{
				index: 2,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
					{
						elem: &dog{id: 1, name: "dog1"},
						ok:   true,
					},
				},
			},
			expected: []*dog{
				{id: 0, name: "dog0"},
				{id: 1, name: "dog1"},
			},
		},
		{
			name: "some removed elements",
			repo: &sliceRepo[*dog]{
				index: 3,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
					{
						elem: &dog{id: 1, name: "dog1"},
						ok:   false,
					},
					{
						elem: &dog{id: 2, name: "dog2"},
						ok:   true,
					},
				},
			},
			expected: []*dog{
				{id: 0, name: "dog0"},
				{id: 2, name: "dog2"},
			},
		},
		{
			name: "all removed elements",
			repo: &sliceRepo[*dog]{
				index: 3,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   false,
					},
					{
						elem: &dog{id: 1, name: "dog1"},
						ok:   false,
					},
					{
						elem: &dog{id: 2, name: "dog2"},
						ok:   false,
					},
				},
			},
			expected: []*dog{},
		},
	}
	testGetAll("sliceRepo[*user]", t, userTests)
	testGetAll("sliceRepo[*dog]", t, dogTests)
}

func TestSliceGetByID(t *testing.T) {
	type test[T IDer] struct {
		name  string
		slice *sliceRepo[T]
		id    int
		err   error
	}
	userTests := []getByIDTest[*user]{
		{
			name: "user slice, id present",
			repo: &sliceRepo[*user]{
				index: 1,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
				},
			},
			id: 0,
		},
		{
			name: "user slice, id not present",
			repo: &sliceRepo[*user]{
				index: 1,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
				},
			},
			id:  1,
			err: errors.New("not found"),
		},
		{
			name: "user slice, id already deleted",
			repo: &sliceRepo[*user]{
				index: 1,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   false,
					},
				},
			},
			id:  0,
			err: errors.New("not found"),
		},
		{
			name: "empty user slice",
			repo: &sliceRepo[*user]{
				index: 0,
				data:  []element[*user]{},
			},
			id:  0,
			err: errors.New("not found"),
		},
	}
	dogTests := []getByIDTest[*dog]{
		{
			name: "dog slice, id present",
			repo: &sliceRepo[*dog]{
				index: 1,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
				},
			},
			id: 0,
		},
		{
			name: "dog slice, id not present",
			repo: &sliceRepo[*dog]{
				index: 1,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
				},
			},
			id:  1,
			err: errors.New("not found"),
		},
		{
			name: "dog slice, id already deleted",
			repo: &sliceRepo[*dog]{
				index: 1,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   false,
					},
				},
			},
			id:  0,
			err: errors.New("not found"),
		},
		{
			name: "empty dog slice",
			repo: &sliceRepo[*dog]{
				index: 0,
				data:  []element[*dog]{},
			},
			id:  0,
			err: errors.New("not found"),
		},
	}
	testGetByID("sliceRepo[*user]", t, userTests)
	testGetByID("sliceRepo[*dog]", t, dogTests)
}

func TestSliceCreate(t *testing.T) {
	userTests := []createTest[*user]{
		{
			name: "empty slice",
			repo: &sliceRepo[*user]{
				index: 0,
				data:  []element[*user]{},
			},
			expected: &sliceRepo[*user]{
				index: 1,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
				},
			},
			item: &user{name: "user0"},
			id:   0,
		},
		{
			name: "some elements slice",
			repo: &sliceRepo[*user]{
				index: 3,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
					{
						elem: &user{id: 1, name: "user1"},
						ok:   true,
					},
					{
						elem: &user{id: 2, name: "user2"},
						ok:   false,
					},
				},
			},
			expected: &sliceRepo[*user]{
				index: 4,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
					{
						elem: &user{id: 1, name: "user1"},
						ok:   true,
					},
					{
						elem: &user{id: 2, name: "user2"},
						ok:   false,
					},
					{
						elem: &user{id: 3, name: "user3"},
						ok:   true,
					},
				},
			},
			item: &user{name: "user3"},
			id:   3,
		},
	}
	dogTests := []createTest[*dog]{
		{
			name: "empty slice",
			repo: &sliceRepo[*dog]{
				index: 0,
				data:  []element[*dog]{},
			},
			expected: &sliceRepo[*dog]{
				index: 1,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
				},
			},
			item: &dog{name: "dog0"},
			id:   0,
		},
		{
			name: "some elements slice",
			repo: &sliceRepo[*dog]{
				index: 3,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
					{
						elem: &dog{id: 1, name: "dog1"},
						ok:   true,
					},
					{
						elem: &dog{id: 2, name: "dog2"},
						ok:   false,
					},
				},
			},
			expected: &sliceRepo[*dog]{
				index: 4,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
					{
						elem: &dog{id: 1, name: "dog1"},
						ok:   true,
					},
					{
						elem: &dog{id: 2, name: "dog2"},
						ok:   false,
					},
					{
						elem: &dog{id: 3, name: "dog3"},
						ok:   true,
					},
				},
			},
			item: &dog{name: "dog3"},
			id:   3,
		},
	}
	testCreate("sliceRepo[*user]", t, userTests)
	testCreate("sliceRepo[*dog]", t, dogTests)
}

func TestSliceRemove(t *testing.T) {
	userTests := []removeTest[*user]{
		{
			name: "empty slice",
			repo: &sliceRepo[*user]{
				index: 0,
				data:  []element[*user]{},
			},
			id:  0,
			err: errors.New("not found"),
		},
		{
			name: "id out of range",
			repo: &sliceRepo[*user]{
				index: 2,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
					{
						elem: &user{id: 1, name: "user1"},
						ok:   true,
					},
				},
			},
			id:  2,
			err: errors.New("not found"),
		},
		{
			name: "element removed",
			repo: &sliceRepo[*user]{
				index: 2,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
					{
						elem: &user{id: 1, name: "user1"},
						ok:   true,
					},
				},
			},
			id:          1,
			removedItem: &user{id: 1, name: "user1"},
			expected: &sliceRepo[*user]{
				index: 2,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
					{
						elem: &user{id: 1, name: "user1"},
						ok:   false,
					},
				},
			},
		},
		{
			name: "element not ok",
			repo: &sliceRepo[*user]{
				index: 2,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
					{
						elem: &user{id: 1, name: "user1"},
						ok:   false,
					},
				},
			},
			id:  1,
			err: errors.New("not found"),
			expected: &sliceRepo[*user]{
				index: 2,
				data: []element[*user]{
					{
						elem: &user{id: 0, name: "user0"},
						ok:   true,
					},
					{
						elem: &user{id: 1, name: "user1"},
						ok:   false,
					},
				},
			},
		},
	}
	dogTests := []removeTest[*dog]{
		{
			name: "empty slice",
			repo: &sliceRepo[*dog]{
				index: 0,
				data:  []element[*dog]{},
			},
			id:  0,
			err: errors.New("not found"),
		},
		{
			name: "id out of range",
			repo: &sliceRepo[*dog]{
				index: 2,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
					{
						elem: &dog{id: 1, name: "dog1"},
						ok:   true,
					},
				},
			},
			id:  2,
			err: errors.New("not found"),
		},
		{
			name: "element removed",
			repo: &sliceRepo[*dog]{
				index: 2,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
					{
						elem: &dog{id: 1, name: "dog1"},
						ok:   true,
					},
				},
			},
			id:          1,
			removedItem: &dog{id: 1, name: "dog1"},
			expected: &sliceRepo[*dog]{
				index: 2,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
					{
						elem: &dog{id: 1, name: "dog1"},
						ok:   false,
					},
				},
			},
		},
		{
			name: "element not ok",
			repo: &sliceRepo[*dog]{
				index: 2,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
					{
						elem: &dog{id: 1, name: "dog1"},
						ok:   false,
					},
				},
			},
			id:  1,
			err: errors.New("not found"),
			expected: &sliceRepo[*dog]{
				index: 2,
				data: []element[*dog]{
					{
						elem: &dog{id: 0, name: "dog0"},
						ok:   true,
					},
					{
						elem: &dog{id: 1, name: "dog1"},
						ok:   false,
					},
				},
			},
		},
	}
	testRemove("sliceRepo[*user]", t, userTests)
	testRemove("sliceRepo[*dog]", t, dogTests)
}
