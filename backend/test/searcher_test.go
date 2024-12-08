package searcher_test

import (
	"testing"

	"challenge-gosolve/backend/app/service"
)

var (
	data []int
	di   service.DataInitializer   = nil
	se   service.SearcherInterface = nil
)

func init() {
	data = mountList(100)
	di = NewDataMock(data)
	se = service.NewSearcher(di)
}

func TestSearcherOneFinder(t *testing.T) {
	key, val := pickRandom(data)
	index, err := se.FindIndex(val)
	if index != key || err != nil {
		t.Fatalf(`index don't match, expected_val: %v, found: %v, error: %v`, key, index, err)
	}
}

func TestSearcherManyFinder(t *testing.T) {
	retry := 25
	for retry > 0 {
		retry--
		key, val := pickRandom(data)
		index, err := se.FindIndex(val)
		if index != key || err != nil {
			t.Fatalf(`index don't match, expected_val: %v, found: %v, error: %v`, key, index, err)
		}
	}
}

func TestSearcherNear(t *testing.T) {
	key := 6
	val := 550
	index, err := se.FindIndex(val)
	if index != key || err != nil {
		t.Fatalf(`index don't match, expected_val: %v, found: %v, error: %v`, key, index, err)
	}
}

func TestSearcherNotNear(t *testing.T) {
	val := 445
	index, err := se.FindIndex(val)
	if err == nil {
		t.Fatalf(`found unexpected index, expected_err: nil, found: %v, error: %v`, index, err)
	}
}

func TestSearcherOutOfRange(t *testing.T) {
	val := 99999
	index, err := se.FindIndex(val)
	if err == nil {
		t.Fatalf(`found unexpected index, expected_err: nil, found: %v, error: %v`, index, err)
	}
}
