package main

import "testing"

func TestCalculateTable(t *testing.T) {
	tests := []struct {
		input      int
		wantResult int
	}{
		{2, 4},
		{6, 8},
		{-2, 0},
		{9999, 10001},
		{3, 5},
	}
	for _, tt := range tests {
		if gotResult := CalculatePlusTwo(tt.input); gotResult != tt.wantResult {
			t.Errorf("CalculatePlusTwo(%v) = %v, want %v", tt.input, gotResult, tt.wantResult)
		}
	}
}
func TestCalculate(t *testing.T) {
	if CalculatePlusTwo(0) != 2 {
		t.Error("Failed")
	}
}
