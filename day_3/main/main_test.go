package main

import (
	mullitover "day_3/Mull_It_Over"
	part_two "day_3/part_two"
	"testing"
)

func TestUncorruptedInstrucions(t *testing.T) {
	uncorruptedMul := mullitover.ScanForUncorruptedInstuctions("../input_test.txt")
	if uncorruptedMul != 161 {
		t.Errorf("Expected add results to be 161, got %d", uncorruptedMul)
	}

}

func TestEnabledUncorruptedInstructions(t *testing.T) {
	enabledMul := part_two.CalculateEnabledMultiplications("../input_test_2.txt")

	if enabledMul != 48 {
		t.Errorf("Expected add results to be 48, got %d", enabledMul)
	}
}
