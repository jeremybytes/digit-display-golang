package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jeremybytes/digit-display-golang/display"
	"github.com/jeremybytes/digit-display-golang/fileloader"
	"github.com/jeremybytes/digit-display-golang/recognize"
	"github.com/jeremybytes/digit-display-golang/shared"
)

func writeOutput(prediction recognize.Prediction) {
	header := fmt.Sprintf("Actual: %v - Prediction: %v\n", prediction.Actual.Actual, prediction.Predicted.Actual)
	image := display.GetImagesAsString(prediction.Actual.Image, prediction.Predicted.Image)
	fmt.Printf("%v%v", header, image)
}

func main() {

	// command line flags
	countPtr := flag.Int("count", 1000, "number of records to identify")
	offsetPtr := flag.Int("offset", 3000, "starting record in data set")
	classPtr := flag.String("class", "euclidean", "classifier calculation type - currently supported: 'manhattan', 'euclidean'")

	flag.Parse()

	fmt.Println("count:", *countPtr)
	fmt.Println("offset:", *offsetPtr)
	fmt.Println("class:", *classPtr)

	fmt.Print("\033[H\033[2J")
	fmt.Printf("STARTING...\n")
	startTime := time.Now()

	training, validation, err := fileloader.LoadData("./data/train.csv", *offsetPtr, *countPtr)
	if err != nil {
		log.Fatalf("Unable to load data: %v", err)
	}

	var classifier recognize.Classifier

	switch *classPtr {
	case "mahattan":
		classifier = &recognize.ManhattanClassifier{}
	case "euclidean":
		classifier = &recognize.EuclideanClassifier{}
	default:
		classifier = &recognize.ManhattanClassifier{}
	}

	fmt.Printf("Using %s\n", classifier)

	classifier.Train(training)

	fmt.Println("Done training...")

	var wg sync.WaitGroup
	ch := make(chan recognize.Prediction)
	missed := make(chan recognize.Prediction, 1000)

	for _, record := range validation {
		wg.Add(1)
		go func(record shared.Record) {
			defer wg.Done()
			prediction, err := recognize.GetPrediction(record, classifier)
			if prediction.Predicted.Actual != record.Actual || err != nil {
				// add to missed
				missed <- prediction
			}
			ch <- prediction
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
		writeOutput(p)
		total++
	}
	elapsed := time.Since(startTime)

	fmt.Println(strings.Repeat("=", 115))
	fmt.Println(classifier)
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
		writeOutput(record)
		fmt.Println(strings.Repeat("-", 115))
	}

	fmt.Println(strings.Repeat("=", 115))
	fmt.Println(classifier)
	fmt.Printf("Total records: %v\n", total)
	fmt.Printf("Time elapsed: %v\n\n", elapsed)
	fmt.Printf("Errors: %v", missedCount)
}
