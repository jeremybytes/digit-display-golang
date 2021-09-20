package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jeremybytes/digit-display-golang/display"
	"github.com/jeremybytes/digit-display-golang/fileloader"
	"github.com/jeremybytes/digit-display-golang/recognize"
)

func writeOutput(prediction int, actual int, pixels []int, closest []int) {
	header := fmt.Sprintf("Actual: %v - Prediction: %v\n", actual, prediction)
	image := display.GetImagesAsString(pixels, closest)
	fmt.Printf("%v%v", header, image)
}

type Prediction struct {
	actual     int
	pixels     []int
	prediction int
	closest    []int
}

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("STARTING...\n")
	startTime := time.Now()

	training, validation, err := fileloader.LoadData("./data/train.csv", 3000, 1000)
	if err != nil {
		log.Fatalf("Unable to load data: %v", err)
	}

	//classifier := &recognize.ManhattanClassifier{}
	classifier := &recognize.EuclideanClassifier{}
	classifier.Train(training)

	fmt.Println("Done training...")

	var wg sync.WaitGroup
	ch := make(chan Prediction)
	missed := make(chan Prediction, 1000)

	for _, record := range validation {
		wg.Add(1)
		go func(record string) {
			defer wg.Done()
			actual, pixels, err := recognize.ParseRecord(record)
			if err != nil {
				// maybe log here later; for now, return from this iteration
				return
			}
			prediction, closest, err := recognize.GetPrediction(pixels, classifier)
			if prediction != actual || err != nil {
				// add to missed
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
		fmt.Printf("\033[0;0H") // moves cursor to top left
		writeOutput(p.prediction, p.actual, p.pixels, p.closest)
		total++
	}
	elapsed := time.Since(startTime)

	fmt.Println(strings.Repeat("=", 115))
	fmt.Printf("Total records: %v\n", total)
	fmt.Printf("Time elapsed: %v\n\n", elapsed)
	fmt.Println("Press ENTER to show errors")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()

	fmt.Println(strings.Repeat("=", 115))
	fmt.Println("   MISSED RECORDS   ")
	fmt.Println(strings.Repeat("=", 115))

	missedCount := 0
	for record := range missed {
		missedCount++
		writeOutput(record.prediction, record.actual, record.pixels, record.closest)
		fmt.Println(strings.Repeat("-", 115))
	}

	fmt.Println(strings.Repeat("=", 115))
	fmt.Printf("Total records: %v\n", total)
	fmt.Printf("Time elapsed: %v\n\n", elapsed)
	fmt.Printf("Errors: %v", missedCount)
}
