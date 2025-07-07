package main

import (
	"gorm.io/gorm"
	"github.com/glebarez/sqlite" // pure Go?
)

type Database struct {
	db *gorm.DB
}

func (re *Database) Open(db_file string) {
	var g gorm.Config
	d, err := gorm.Open(sqlite.Open(db_file), &g)
	if err != nil {
		panic(err.Error())
	}
	re.db = d
}

func (re Database) Init(sk StateKeeper) {
	re.db.AutoMigrate(sk)
}

/*
func (re Database) Save(sk StateKeeper) {
	re.db.Create(&sk)
}
*/

type StateKeeper interface {
	Init(Database)
	Save(Database)
}

