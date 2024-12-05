package service

import (
	"errors"
)

type DataType []int

type database struct {
	data DataType
}

type DatabaseInterface interface {
	FindIndex(int) (int, error)
}

func (db *database) FindIndex(index int) (int, error) {
	if 0 <= index && index < len(db.data) {
		return db.data[index], nil;
	}
	return 0, errors.New("not_found")
}

func NewDatabase() DatabaseInterface {
	var db = new(database)
	db.data = []int{ 0, 100, 200, 300, 400, 500, 600, 700, 800, 900 }
	return db
}
