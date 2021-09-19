package display

import "fmt"

func OutputImage(imageData []int) {
	// build a string
	var output string
	for i, pixel := range imageData {
		outputChar := getDisplayCharForPixel(pixel)
		output += outputChar
		output += outputChar
		if i%28 == 0 {
			output += "\n"
		}
	}
	output += "\n"
	fmt.Print(output)
}

func getDisplayCharForPixel(i int) string {
	if i > 16 && i < 32 {
		return "."
	}
	if i >= 32 && i < 64 {
		return ":"
	}
	if i >= 64 && i < 160 {
		return "o"
	}
	if i >= 160 && i < 224 {
		return "O"
	}
	if i >= 224 {
		return "@"
	}
	return " "
}
