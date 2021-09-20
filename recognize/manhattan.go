package recognize

import "fmt"

type ManhattanClassifier struct {
	TrainingData []Observation
}

func (c *ManhattanClassifier) Train(traingingData []string) error {
	for _, record := range traingingData {
		actual, err := stringToActual(record)
		if err != nil {
			// maybe log error here
			// for now, skip this record
			continue
		}
		pixels, err := stringToIntArray(record)
		if err != nil {
			// maybe log error here
			// for now, skip this record
			continue
		}
		obs := Observation{actual, pixels}
		c.TrainingData = append(c.TrainingData, obs)
	}
	if len(c.TrainingData) <= 0 {
		return fmt.Errorf("Train produced no valid traning data")
	}
	return nil
}

func Abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func (c ManhattanClassifier) Predict(pixels []int) (prediction int, closest []int, err error) {
	bestPrediction := -1
	var bestPixels []int = nil
	var bestTotal int = 100000000
	for _, train := range c.TrainingData {
		total := 0
		for i := range pixels {
			total = total + Abs(pixels[i]-train.pixels[i])
		}
		if total < bestTotal {
			bestTotal = total
			bestPrediction = train.actual
			bestPixels = train.pixels
		}
	}

	if bestPrediction == -1 {
		return -1, nil, fmt.Errorf("Unable to get a valid prediction")
	}

	return bestPrediction, bestPixels, err
}
