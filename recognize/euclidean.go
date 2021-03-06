package recognize

import (
	"fmt"
)

type EuclideanClassifier struct {
	TrainingData []Record
}

func (c EuclideanClassifier) String() string {
	return "Euclidean Classifier"
}

func (c *EuclideanClassifier) Train(traingingData []Record) error {
	c.TrainingData = traingingData
	return nil
}

func (c EuclideanClassifier) Predict(input Record) (prediction Prediction, err error) {
	best := Record{Actual: -1, Image: nil}
	var bestTotal int = 100000000
	for _, train := range c.TrainingData {
		total := 0
		for i := range input.Image {
			diff := input.Image[i] - train.Image[i]
			total = total + (diff * diff)
		}
		if total < bestTotal {
			bestTotal = total
			best = train
		}
	}

	if best.Actual == -1 {
		return Prediction{}, fmt.Errorf("unable to get a valid prediction")
	}

	return Prediction{Actual: input, Predicted: best}, err
}
