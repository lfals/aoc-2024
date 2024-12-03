package parttwo

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func CalculateEnabledMultiplications(inputFilePath string) int {
	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	line := strings.ReplaceAll(string(content), "\n", "")

	dontReplaced := strings.ReplaceAll(line, "don't()", "\ndon't()")
	doReplaced := strings.ReplaceAll(dontReplaced, "do()", "\ndo()")

	splitted := strings.Split(doReplaced, "\n")

	uncorruptedInstruction := 0
	regex := `mul\(\d{1,3},\d{1,3}\)`

	numRegex := `\d{1,3},\d{1,3}`

	for i := range splitted {
		if strings.HasPrefix(splitted[i], "don't()") == false {
			re := regexp.MustCompile(regex)
			matches := re.FindAllString(splitted[i], -1)

			for j := range matches {
				numre := regexp.MustCompile(numRegex)
				nummatches := numre.FindAllString(matches[j], -1)
				nums := strings.Split(nummatches[0], ",")
				left, _ := strconv.Atoi(nums[0])
				right, _ := strconv.Atoi(nums[1])
				mult := left * right
				uncorruptedInstruction += mult

			}
		}
	}

	return uncorruptedInstruction
}
