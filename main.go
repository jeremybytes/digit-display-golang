package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jeremybytes/digit-display-golang/display"
	"github.com/jeremybytes/digit-display-golang/fileloader"
	"github.com/jeremybytes/digit-display-golang/recognize"
)

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

func writeOutput(prediction int, actual int, pixels []int) {
	fmt.Printf("Actual: %v - Prediction: %v", actual, prediction)
	display.OutputImage(pixels)
}

type Prediction struct {
	actual     int
	pixels     []int
	prediction int
	closest    []int
}

func main() {
	fmt.Printf("STARTING...\n")
	startTime := time.Now()

	training, validation, err := fileloader.LoadData("./data/train.csv", 3000, 1000)
	if err != nil {
		log.Fatalf("%v", err)
	}

	manClassifier := &recognize.ManhattanClassifier{}
	manClassifier.Train(training)

	eucClassifier := &recognize.EuclideanClassifier{}
	eucClassifier.Train(training)

	fmt.Println("Done training...")

	var wg sync.WaitGroup
	ch := make(chan Prediction)
	missed := make(chan Prediction, 1000)

	for _, record := range validation {
		wg.Add(1)
		go func(record string) {
			defer wg.Done()
			actual, _ := stringToActual(record)
			pixels, _ := stringToIntArray(record)
			prediction, closest := recognize.GetPrediction(pixels, eucClassifier)
			if prediction != actual {
				missed <- Prediction{actual, pixels, prediction, closest}
			}
			ch <- Prediction{actual, pixels, prediction, closest}
		}(record)
	}

	go func() {
		wg.Wait()
		close(ch)
		close(missed)
	}()

	total := 0
	for p := range ch {
		writeOutput(p.prediction, p.actual, p.pixels)
		total++
	}
	elapsed := time.Since(startTime)

	fmt.Println("==============================================")
	fmt.Printf("Total records: %v\n", total)
	fmt.Printf("Time elapsed: %v\n\n", elapsed)
	fmt.Println("Press ENTER to show errors")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()

	fmt.Println("==============================================")
	fmt.Println("   MISSED RECORDS   ")
	fmt.Println("==============================================")

	missedCount := 0
	for record := range missed {
		missedCount++
		writeOutput(record.prediction, record.actual, record.pixels)
		display.OutputImage(record.closest)
	}

	fmt.Println("==============================================")
	fmt.Printf("Total records: %v\n", total)
	fmt.Printf("Time elapsed: %v\n\n", elapsed)
	fmt.Printf("Errors: %v", missedCount)
}
