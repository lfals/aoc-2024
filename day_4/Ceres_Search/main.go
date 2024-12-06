package ceressearch

import (
	"log"
	"os"
	"strings"
)

// file, err := os.Open(inputFilePath)
// if err != nil {
// 	log.Fatalf("failed to read file: %v", err)
// }

// defer file.Close()

// scanner := bufio.NewScanner(file)

// for scanner.Scan() {
// 	line := scanner.Text()

//		fmt.Println(line)
//	}
func CountXmas(inputFilePath string) int {
	totalXmas := 0

	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	for lineIndex, line := range lines {
		chars := strings.Split(line, "")
		for charIndex, char := range chars {

			//Por linha
			if charIndex < len(chars)-3 {
				text := char + strings.Join(chars[charIndex+1:charIndex+4], "")
				if text == "XMAS" || text == "SAMX" {
					totalXmas += 1
				}
			}

			// Por diagonal direita
			if charIndex < len(chars)-3 {
				if lineIndex < len(lines)-3 {
					text := char + lines[lineIndex+1][charIndex+1:charIndex+2] + lines[lineIndex+2][charIndex+2:charIndex+3] + lines[lineIndex+3][charIndex+3:charIndex+4]
					if text == "XMAS" || text == "SAMX" {

						totalXmas += 1
					}
				}
			}

			// Por vertical
			if lineIndex < len(lines)-3 {
				text := char + lines[lineIndex+1][charIndex:charIndex+1] + lines[lineIndex+2][charIndex:charIndex+1] + lines[lineIndex+3][charIndex:charIndex+1]
				if text == "XMAS" || text == "SAMX" {

					totalXmas += 1
				}
			}

			// Por diagonal esquerda
			if charIndex > 2 {
				if lineIndex < len(lines)-3 {
					text := char + lines[lineIndex+1][charIndex-1:charIndex] + lines[lineIndex+2][charIndex-2:charIndex-1] + lines[lineIndex+3][charIndex-3:charIndex-2]
					if text == "XMAS" || text == "SAMX" {

						totalXmas += 1
					}
				}
			}

		}
	}

	return totalXmas
}
