package service

import (
	"errors"
	"sort"
)

type Searcher struct {
	data []int
}

type DataInitializer interface {
	Load() []int
}

type SearcherInterface interface {
	FindIndex(int) (int, error)
}

func (se *Searcher) isNear(found int, value int) bool {
	percent := float64(found) / float64(value)
	return 0.9 <= percent && percent <= 1.1
}

func (se *Searcher) FindIndex(value int) (int, error) {
	length := len(se.data)
	index := sort.Search(length, func(index int) bool {
		return se.data[index] >= value
	})
	if index < 0 || length <= index {
		return 0, errors.New("not_found")
	}
	var valuFound = se.data[index]
	if valuFound == value {
		return index, nil
	}
	if se.isNear(valuFound, value) {
		return index, nil
	}
	return 0, errors.New("not_found")
}

func NewSearcher(di DataInitializer) SearcherInterface {
	var se = new(Searcher)
	se.data = di.Load()
	return se
}
