package historian_hysteria

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalculateDistance(file string) int {

	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	left := []int{}
	right := []int{}
	totalDistance := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")

		leftInt, _ := strconv.Atoi(parts[0])
		rightInt, _ := strconv.Atoi(parts[len(parts)-1])

		left = append(left, leftInt)
		right = append(right, rightInt)

		sort.Ints(left)
		sort.Ints(right)

	}

	for i := 0; i < len(left); i++ {
		totalDistance += left[i] - right[i]
	}

	fmt.Println(totalDistance)

	return int(math.Abs(float64(totalDistance)))
}
