package main

import (
	ceressearch "day_4/Ceres_Search"
	"testing"
)

func TestUncorruptedInstrucions(t *testing.T) {
	totalXmas := ceressearch.CountXmas("../input_test.txt")
	if totalXmas != 18 {
		t.Errorf("Expected add results to be 18, got %d", totalXmas)
	}

}
