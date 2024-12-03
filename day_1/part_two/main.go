package part_two

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func CalculateSimilarityScore(inputFilePath string) int {
	similarityScore := 0

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := []string{}
	right := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		left = append(left, parts[0])
		right = append(right, parts[len(parts)-1])

	}

	for i := 0; i < len(left); i++ {
		count := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				count += 1
			}
		}

		leftInt, _ := strconv.Atoi(left[i])

		similarityScore += count * leftInt

	}

	return similarityScore
}
