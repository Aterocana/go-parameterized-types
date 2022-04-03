package main

import "fmt"

// dog can be a type contained in a repo
type dog struct {
	id    int
	name  string
	breed string
}

func (d *dog) GetID() int {
	return d.id
}

func (d *dog) SetID(id int) {
	d.id = id
}

func (d *dog) String() string {
	return fmt.Sprintf("<#%d %s (%s)>", d.id, d.name, d.breed)
}
