package service

import (
	"log"
	"os"
	"bufio"
	"strconv"
	"sort"
	"errors"
)

type DataType []int

type Reader struct {
	filepath string
	data DataType
}

type ReaderInterface interface {
	Init()
	Search(int) (int, error)
}

func (re *Reader) Init() {
	file, err := os.Open(re.filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var data []int
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("[Reader] file %v loaded", len(data))
	re.data = data
}

func (re *Reader) Search(value int) (int, error) {
	length := len(re.data)
	index := sort.Search(length, func(index int) bool {
		return re.data[index] >= value
	})
	if index < 0 || length <= index {
		return 0, errors.New("not_found")
	}
	var found = re.data[index]
	if found == value {
		return index, nil
	}
	if float64(value) * 0.9 <= float64(found) && float64(found) <= float64(value) * 1.1 {
		return index, nil
	}
	return 0, errors.New("not_found")
}

func NewReader(filepath string) ReaderInterface {
	var re = new(Reader)
	re.filepath = filepath
	return re
}
