package searcher_test

import (
	"math/rand/v2"
)

func mountList(size int) []int {
	list := make([]int, size)
	val := 0
	for key, _ := range list {
		list[key] = val
		val += 100
	}
	return list
}

func pickRandom(list []int) (int, int) {
	size := len(list)
	key := rand.IntN(size)
	val := list[key]
	return key, val
}
