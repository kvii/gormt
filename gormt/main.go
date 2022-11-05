package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"))
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	db.Exec("SELECT 1")
}
