package recognize

type EuclideanClassifier struct {
	TrainingData []Observation
}

func (c *EuclideanClassifier) Train(traingingData []string) {
	for _, record := range traingingData {
		actual, _ := stringToActual(record)
		pixels, _ := stringToIntArray(record)
		obs := Observation{actual, pixels}
		c.TrainingData = append(c.TrainingData, obs)
	}
}

func (c EuclideanClassifier) Predict(pixels []int) (prediction int, closest []int) {
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
	return bestPrediction, bestPixels
}
