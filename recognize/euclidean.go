package recognize

import "fmt"

type EuclideanClassifier struct {
	TrainingData []Observation
}

func (c *EuclideanClassifier) Train(traingingData []string) error {
	for _, record := range traingingData {
		actual, pixels, err := ParseRecord(record)
		if err != nil {
			// maybe log here later; for now, return from this iteration
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

func (c EuclideanClassifier) Predict(pixels []int) (prediction int, closest []int, err error) {
	bestPrediction := -1
	var bestPixels []int = nil
	var bestTotal int = 100000000
	for _, train := range c.TrainingData {
		total := 0
		for i := range pixels {
			diff := pixels[i] - train.pixels[i]
			total = total + (diff * diff)
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

	return bestPrediction, bestPixels, nil
}
