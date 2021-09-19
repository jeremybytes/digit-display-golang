package recognize

import (
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
	output, _ := strconv.Atoi(items[0])
	return output, nil
}

func stringToIntArray(record string) ([]int, error) {
	items := strings.Split(record, ",")
	ints := make([]int, 784)
	for i, pixel := range items[1:] {
		output, _ := strconv.Atoi(pixel)
		ints[i] = output
	}
	return ints, nil
}

func GetPrediction(pixels []int, classifier Classifier) (prediction int, closest []int) {
	prediction, closest = classifier.Predict(pixels)
	return
}
