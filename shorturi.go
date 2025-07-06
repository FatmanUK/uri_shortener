package main

import (
	"gorm.io/gorm"
	"github.com/glebarez/sqlite" // pure Go?
)

type Database struct {
	db *gorm.DB
}

type ShortUri struct {
	gorm.Model
	Id    uint
	Nick  string
	Uri   string
}

func (re *ShortUri) OpenDb(dbfile string) Database {
	d, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	d.AutoMigrate(&ShortUri{})
	db := Database{d}
	return db
}

func (re *ShortUri) FindById(d Database, id uint) error {
	result := d.db.Where("id = ?", id).Last(re)
	return result.Error
}

func (re *ShortUri) FindByNick(d Database, nick string) error {
	result := d.db.Where("nick = ?", nick).Last(re)
	return result.Error
}

func (re *ShortUri) Create(d Database) {
	d.db.Create(re)
}

