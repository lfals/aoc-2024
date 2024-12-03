package mullitover

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ScanForUncorruptedInstuctions(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	defer file.Close()

	uncorruptedInstruction := 0

	scanner := bufio.NewScanner(file)

	regex := `mul\(\d{1,3},\d{1,3}\)`

	numRegex := `\d{1,3},\d{1,3}`

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(regex)
		matches := re.FindAllString(line, -1)

		for i := range matches {
			numre := regexp.MustCompile(numRegex)
			nummatches := numre.FindAllString(matches[i], -1)
			nums := strings.Split(nummatches[0], ",")
			left, _ := strconv.Atoi(nums[0])
			right, _ := strconv.Atoi(nums[1])
			mult := left * right

			uncorruptedInstruction += mult
		}
	}

	return uncorruptedInstruction
}
