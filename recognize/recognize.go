package recognize

import (
	"fmt"
	"strconv"
	"strings"
)

type Observation struct {
	actual int
	pixels []int
}

type Classifier interface {
	Train(trainingData []string)
	Predict(pixels []int) (prediction int, closest []int)
}

func stringToActual(record string) (int, error) {
	items := strings.Split(record, ",")
	output, err := strconv.Atoi(items[0])
	if err != nil {
		return -1, fmt.Errorf("Unable to parse actual value: %v", err)
	}
	return output, nil
}

func stringToIntArray(record string) ([]int, error) {
	items := strings.Split(record, ",")
	ints := make([]int, 784)
	for i, pixel := range items[1:] {
		output, err := strconv.Atoi(pixel)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse pixel value (%s): %v", pixel, err)
		}
		ints[i] = output
	}
	return ints, nil
}

func GetPrediction(pixels []int, classifier Classifier) (prediction int, closest []int) {
	prediction, closest = classifier.Predict(pixels)
	return
}
