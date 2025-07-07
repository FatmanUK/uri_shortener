package main

// go test database.go database_test.go -v

import (
	"testing"
)

type TestThing struct {
	Id uint
	Name string
}

func (re TestThing) Init(d Database) {
	d.Init(re)
}

func (re TestThing) Save(d Database) {
	//d.Save(re)
	d.db.Create(&re)
}

/*
func (re *TestThing) Load(Database) {
}
*/

func TestOpen(t *testing.T) {
	var d Database
	d.Open("test.sqlite")
	var r TestThing
	r.Init(d)
}

func TestSave(t *testing.T) {
	var d Database
	d.Open("test.sqlite")
	r1 := TestThing{1, "bob"}
	r1.Save(d)
	r2 := TestThing{2, "greg"}
	r2.Save(d)
}

func TestLoad(t *testing.T) {
}

