package main

import (
	"day_2/Red_Nosed_reports"
	"fmt"
)

func main() {
	safeReports := Red_Nosed_reports.CalculateSafeReports("../input.txt")
	fmt.Println("Part One:", safeReports)
}
