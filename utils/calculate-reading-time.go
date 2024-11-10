package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculateWordCount() int {
	articleFile, err := os.Open("data/articles/file-integrity.txt")
	if err != nil {
		log.Println("Error opening file:", err)
		return 0
	}
	defer articleFile.Close()

	wordCount := 0
	scanner := bufio.NewScanner(articleFile)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	return wordCount
}

func CalculateReadingTime() int {
	// Assuming the average speed is 200 words per minute
	wordCount := calculateWordCount()
	return (wordCount + 199) / 200
}
