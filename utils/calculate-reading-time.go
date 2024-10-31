package utils

import "strings"

func CalculateReadingTime(content string) int {
	// Assuming the average speed is 200 words per minute
	wordCount := len(strings.Fields(content))
	return (wordCount + 199) / 200
}
