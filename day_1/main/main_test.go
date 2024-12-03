package main

import (
	"day_1/historian_hysteria"
	"day_1/part_two"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	distance := historian_hysteria.CalculateDistance("../input_test.txt")
	if distance != 11 {
		t.Errorf("Expected distance to be 11, got %d", distance)
	}
}

func TestCalculateSimilarityScore(t *testing.T) {
	similarityScore := part_two.CalculateSimilarityScore("../input_test.txt")
	if similarityScore != 31 {
		t.Errorf("Expected similarity score to be 31, got %d", similarityScore)
	}
}
