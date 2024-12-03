package main_test

import (
	"day_2/Red_Nosed_reports"
	parttwo "day_2/part_two"
	"testing"
)

func TestCalculateSafeReports(t *testing.T) {
	safeReports := Red_Nosed_reports.CalculateSafeReports("../input_test.txt")
	if safeReports != 2 {
		t.Errorf("Expected safe reports to be 2, got %d", safeReports)
	}
}

func TestCalculateSafeReportsDamper(t *testing.T) {
	safeReports := parttwo.ProblemDampener("../input_test.txt")
	if safeReports != 4 {
		t.Errorf("Expected safe reports to be 4, got %d", safeReports)
	}
}
