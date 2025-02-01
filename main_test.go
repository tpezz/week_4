package main

import (
	"testing"
)

// TestReadCSV ensures that readCSV function runs without errors.
func TestReadCSV(t *testing.T) {
	filePath := "Anscombe_quartet_data.csv"

	// Call the function
	xVals, yVals := readCSV(filePath)

	// Check if data was read (basic sanity check)
	if len(xVals) == 0 || len(yVals) == 0 {
		t.Errorf("readCSV returned empty slices")
	}
}
