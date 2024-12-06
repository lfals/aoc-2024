package main

import (
	ceressearch "day_4/Ceres_Search"
	part_two "day_4/part_two"
	"testing"
)

func TestXmas(t *testing.T) {
	totalXmas := ceressearch.CountXmas("../input_test.txt")
	if totalXmas != 18 {
		t.Errorf("Expected add results to be 18, got %d", totalXmas)
	}

}

func TestX_MAS(t *testing.T) {
	totalX_mas := part_two.CalculateX_MAS("../input_test.txt")
	if totalX_mas != 9 {
		t.Errorf("Expected add results to be 9, got %d", totalX_mas)
	}
}
