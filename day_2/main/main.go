package main

import (
	"day_2/Red_Nosed_reports"
	parttwo "day_2/part_two"
	"fmt"
)

func main() {
	safeReports := Red_Nosed_reports.CalculateSafeReports("../input.txt")
	fmt.Println("Part One:", safeReports)

	safeReportsDamper := parttwo.ProblemDampener("../input.txt")
	fmt.Println("Part Two:", safeReportsDamper)
}
