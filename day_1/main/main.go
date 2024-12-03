package main

import (
	"day_1/part_two"
	"fmt"
)

func main() {
	// distance := historian_hysteria.CalculateDistance("../input.txt")
	// fmt.Println("Part One:", distance)

	similarityScore := part_two.CalculateSimilarityScore("../input.txt")
	fmt.Println("Part Two:", similarityScore)
}
