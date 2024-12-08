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
	var valueFound = se.data[index]
	if valueFound == value {
		return index, nil
	}
	if se.isNear(valueFound, value) {
		return index, nil
	}
	index -= 1
	if index < 0 {
		return 0, errors.New("not_found")
	}
	var beforeValue = se.data[index]
	if se.isNear(beforeValue, value) {
		return index, nil
	}
	return 0, errors.New("not_found")
}

func NewSearcher(di DataInitializer) SearcherInterface {
	var se = new(Searcher)
	se.data = di.Load()
	return se
}
