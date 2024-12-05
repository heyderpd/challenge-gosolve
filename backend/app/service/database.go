package service

import (
	"errors"
	"sort"
)

type DataType []int

type database struct {
	data DataType
}

type DatabaseInterface interface {
	FindIndex(int) (int, error)
}

func (db *database) FindIndex(value int) (int, error) {
	length := len(db.data)
	index := sort.Search(length, func(index int) bool {
		return db.data[index] >= value
	})
	if index < 0 || length <= index {
		return 0, errors.New("not_found")
	}
	var found = db.data[index]
	if found == value {
		return index, nil
	}
	if float64(value) * 0.9 <= float64(found) && float64(found) <= float64(value) * 1.1 {
		return index, nil
	}
	return 0, errors.New("not_found")
}

func NewDatabase(re ReaderInterface) DatabaseInterface {
	var db = new(database)
	db.data = re.Load()
	return db
}
