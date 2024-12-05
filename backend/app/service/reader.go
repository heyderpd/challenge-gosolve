package service

import (
	"log"
	"os"
	"bufio"
	"strconv"
)

type Reader struct {
	filepath string
}

type ReaderInterface interface {
	Load() []int
}

func (re *Reader) Load() []int {
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
	return data
}

func NewReader(filepath string) ReaderInterface {
	var re = new(Reader)
	re.filepath = filepath
	return re
}
