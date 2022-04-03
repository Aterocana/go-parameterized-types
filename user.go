package main

import "fmt"

// user can be a type contained in a repo
type user struct {
	id   int
	name string
}

func (u *user) GetID() int {
	return u.id
}

func (u *user) SetID(id int) {
	u.id = id
}

func (u *user) String() string {
	return fmt.Sprintf("<#%d %s>", u.id, u.name)
}
