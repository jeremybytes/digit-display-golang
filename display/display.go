package display

import "fmt"

func OutputImage(imageData []int) {
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
	switch {
	case (i > 16 && i < 32):
		return "."
	case (i >= 32 && i < 64):
		return ":"
	case (i >= 64 && i < 160):
		return "o"
	case (i >= 160 && i < 224):
		return "O"
	case (i >= 224):
		return "@"
	default:
		return " "
	}
}
