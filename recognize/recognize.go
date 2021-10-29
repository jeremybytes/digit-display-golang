package recognize

import (
	"fmt"

	"github.com/jeremybytes/digit-display-golang/shared"
)

type Prediction struct {
	Actual    shared.Record
	Predicted shared.Record
}

type Classifier interface {
	Train(trainingData []shared.Record) error
	Predict(input shared.Record) (prediction Prediction, err error)
}

func GetPrediction(input shared.Record, classifier Classifier) (prediction Prediction, err error) {
	prediction, err = classifier.Predict(input)
	if err != nil {
		return Prediction{}, fmt.Errorf("method Predict failed: %v", err)
	}
	return
}
