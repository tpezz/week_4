package main

import (
	"testing"
)


func TestComputeRegression(t *testing.T) {
	xValues := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	yValues := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	expectedSlope := 0.5
	expectedIntercept := 3.0

	slope, intercept, _, err := ComputeRegression(xValues, yValues)
	if err != nil {
		t.Fatalf("Failed to compute regression: %v", err)
	}

	if slope < expectedSlope-0.1 || slope > expectedSlope+0.1 {
		t.Errorf("Expected slope ~%.5f, got %.5f", expectedSlope, slope)
	}
	if intercept < expectedIntercept-0.1 || intercept > expectedIntercept+0.1 {
		t.Errorf("Expected intercept ~%.5f, got %.5f", expectedIntercept, intercept)
	}
}