package service

import (
)

type database struct {
	data ReaderInterface
}

type DatabaseInterface interface {
	FindIndex(int) (int, error)
}

func (db *database) FindIndex(value int) (int, error) {
	found, err := db.data.Search(value)
	if err != nil {
		return 0, err
	}
	return found, nil
}

func NewDatabase(re ReaderInterface) DatabaseInterface {
	var db = new(database)
	db.data = re
	re.Init()
	return db
}
