package fileloader

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jeremybytes/digit-display-golang/shared"
)

func trimWhiteSpace(data []string) []string {
	for i, n := range data {
		data[i] = strings.TrimSpace(n)
	}
	return data
}

func LoadData(path string, offset int, recordCount int) (training []shared.Record, validation []shared.Record, err error) {
	dataLines, err := getRawData(path)
	if err != nil {
		return nil, nil, fmt.Errorf("getDataBytes failed: %v", err)
	}
	var allRecords []shared.Record
	for _, line := range dataLines {
		parsed, err := parseRawData(line)
		if err != nil {
			continue
		}
		parsedRecord, err := parseRecord(parsed)
		if err != nil {
			continue
		}
		allRecords = append(allRecords, parsedRecord)
	}

	if (offset + recordCount) > len(dataLines) {
		return nil, nil, fmt.Errorf("LoadPath offset + recordCount is bigger than the available dataset")
	}
	training, validation = splitDataSets(allRecords, offset, recordCount)
	return training, validation, nil
}

func getRawData(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	dataLines := strings.Split(string(bytes), "\n")
	dataLines = trimWhiteSpace(dataLines)
	return dataLines, nil
}

func parseRawData(rawData string) ([]int, error) {
	ints := make([]int, 785)
	items := strings.Split(rawData, ",")
	for i, val := range items {
		output, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("unable to parse integer value (%s): %v", val, err)
		}
		ints[i] = output
	}
	return ints, nil
}

func parseRecord(data []int) (shared.Record, error) {
	if len(data) != 785 {
		return shared.Record{}, fmt.Errorf("incorrect data size; should be 785 found %v", len(data))
	}
	actual := data[0]
	image := data[1:]

	return shared.Record{Actual: actual, Image: image}, nil
}

func splitDataSets(data []shared.Record, offset int, recordCount int) ([]shared.Record, []shared.Record) {
	trainingData := append(data[(1+offset+recordCount+1):], data[1:offset]...)
	validationData := data[(1 + offset):(1 + offset + recordCount)]

	return trainingData, validationData
}
