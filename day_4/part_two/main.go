package parttwo

import (
	"log"
	"os"
	"strings"
)

func CalculateX_MAS(inputFilePath string) int {
	totalXmas := 0

	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	// options := []string{"SSAMM", "MMASS", "SMAMS", "MSASM"}

	for lineIndex, line := range lines {
		chars := strings.Split(line, "")
		for charIndex, char := range chars {
			if charIndex < len(chars)-2 && lineIndex < len(lines)-2 {
				text :=
					char + chars[charIndex+2 : charIndex+3][0] +
						lines[lineIndex+1][charIndex+1:charIndex+2] +
						lines[lineIndex+2][charIndex:charIndex+1] + lines[lineIndex+2][charIndex+2:charIndex+3]

				if text == "MMASS" || text == "MSAMS" || text == "SSAMM" || text == "SMASM" {
					totalXmas += 1

				}
			}

		}
	}

	return totalXmas

}
