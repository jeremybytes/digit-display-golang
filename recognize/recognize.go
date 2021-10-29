package recognize

import (
	"fmt"

	"github.com/jeremybytes/digit-display-golang/shared"
)

type Record = shared.Record

type Prediction struct {
	Actual    Record
	Predicted Record
}

type Classifier interface {
	Train(trainingData []Record) error
	Predict(input Record) (prediction Prediction, err error)
}

func GetPrediction(input Record, classifier Classifier) (prediction Prediction, err error) {
	prediction, err = classifier.Predict(input)
	if err != nil {
		return Prediction{}, fmt.Errorf("method Predict failed: %v", err)
	}
	return
}
