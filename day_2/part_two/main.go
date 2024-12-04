package parttwo

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func ProblemDampener(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReports := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		unsafeReports := 0
		increasing := 0
		decreasing := 0
		for i := 0; i < len(parts)-1; i++ {
			left, _ := strconv.Atoi(parts[i])
			right, _ := strconv.Atoi(parts[i+1])
			diff := float64(left) - float64(right)
			positiveDiff := math.Abs(diff)
			if positiveDiff > 3 || positiveDiff < 1 {
				unsafeReports += 1
			}

			if math.Signbit(diff) {
				increasing += 1
			} else {
				decreasing += 1
			}

		}

		allIncreasing := increasing > 0 && decreasing == 0
		allDecreasing := decreasing > 0 && increasing == 0

		if unsafeReports == 0 && (allIncreasing == true || allDecreasing == true) {
			safeReports += 1
		} else {

		out:
			for k := range parts {

				copy := make([]string, len(parts))
				copy = append(copy[:0], parts...)
				removed := append(copy[:k], copy[k+1:]...)

				unsafeReportsCount := 0
				unsafeIncreasing := 0
				unsafeDecreasing := 0
				for index := 0; index < len(removed)-1; index++ {

					left, _ := strconv.Atoi(removed[index])
					right, _ := strconv.Atoi(removed[index+1])

					diff := float64(left) - float64(right)

					positiveUnsafeDiff := math.Abs(diff)

					if positiveUnsafeDiff > 3 || positiveUnsafeDiff < 1 {
						unsafeReportsCount += 1
					}

					if math.Signbit(diff) {
						unsafeIncreasing += 1
					} else {
						unsafeDecreasing += 1
					}

				}

				allIncreasing := unsafeIncreasing > 0 && unsafeDecreasing == 0
				allDecreasing := unsafeDecreasing > 0 && unsafeIncreasing == 0

				if unsafeReportsCount == 0 && (allIncreasing == true || allDecreasing == true) {
					safeReports += 1
					break out
				}

			}
		}

	}
	return safeReports
}
