package display

import "strings"

func GetImagesAsString(image1 []int, image2 []int) string {
	firstImage := strings.Split(GetImageAsString(image1), "\n")
	secondImage := strings.Split(GetImageAsString(image2), "\n")

	var output string
	// merge strings
	for i := 0; i < 28; i++ {
		output += firstImage[i]
		output += " | "
		output += secondImage[i]
		output += "\n"
	}
	return output
}

func GetImageAsString(imageData []int) string {
	var output string
	for i, pixel := range imageData {
		if i%28 == 0 && i > 0 {
			output += "\n"
		}
		outputChar := getDisplayCharForPixel(pixel)
		output += outputChar
		output += outputChar
	}
	output += "\n"
	return output
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
