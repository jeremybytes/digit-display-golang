package recognize

type ManhattanClassifier struct {
	TrainingData []Observation
}

func (c *ManhattanClassifier) Train(traingingData []string) {
	for _, record := range traingingData {
		actual, _ := stringToActual(record)
		pixels, _ := stringToIntArray(record)
		obs := Observation{actual, pixels}
		c.TrainingData = append(c.TrainingData, obs)
	}
}

func Abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func (c ManhattanClassifier) Predict(pixels []int) (prediction int, closest []int) {
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
	return bestPrediction, bestPixels
}
