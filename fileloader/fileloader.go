package fileloader

import (
	"fmt"
	"os"
	"strings"
)

func trimWhiteSpace(data []string) []string {
	for i, n := range data {
		data[i] = strings.TrimSpace(n)
	}
	return data
}

func LoadData(path string, offset int, recordCount int) (training []string, validation []string, err error) {
	bytes, err := getDataBytes(path)
	if err != nil {
		return nil, nil, fmt.Errorf("getDataBytes failed: %v", err)
	}
	allData := string(bytes)
	dataLines := strings.Split(allData, "\n")

	dataLines = trimWhiteSpace(dataLines)

	if (offset + recordCount) > len(dataLines) {
		return nil, nil, fmt.Errorf("LoadPath offset + recordCount is bigger than the available dataset")
	}

	trainingData := append(dataLines[(1+offset+recordCount+1):], dataLines[1:offset]...)
	validationData := dataLines[(1 + offset):(1 + offset + recordCount)]

	return trainingData, validationData, nil
}

func getDataBytes(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return bytes, nil
}
