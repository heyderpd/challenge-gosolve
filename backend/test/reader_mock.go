package searcher_test

import (
	"challenge-gosolve/backend/app/service"
)

type DataMock struct {
	data []int
}

func (re *DataMock) Load() []int {
	return re.data
}

func NewDataMock(data []int) service.DataInitializer {
	var fr = new(DataMock)
	fr.data = data
	return fr
}
