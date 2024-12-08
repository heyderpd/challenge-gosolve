package service

import (
	"os"
	"bufio"
	"strconv"

	"challenge-gosolve/backend/app/utils"
)

type FileReader struct {
	filepath string
	data []int
}

func (re *FileReader) Load() []int {
	file, err := os.Open(re.filepath)
	defer file.Close()
	if err != nil {
		utils.Log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var data []int
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			utils.Log.Fatal(err)
		}
		data = append(data, line)
	}
	if err := scanner.Err(); err != nil {
		utils.Log.Fatal(err)
	}
	utils.Log.Printf("[Reader] file %v loaded", len(data))
	return data
}

func NewReader(filepath string) DataInitializer {
	var fr = new(FileReader)
	fr.filepath = filepath
	return fr
}
