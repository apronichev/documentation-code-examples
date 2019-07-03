package divisionMod

import (
	"testing"
)

func TestMultiplyTable(t *testing.T) {
	tests := []struct {
		input      int
		wantResult int
	}{
		{4, 2},
		{12, 6},
		{-4, -2},
		{19998, 9999},
		{6, 3},
	}
	for _, tt := range tests {
		if gotResult := CalculateDivisionTwo(tt.input); gotResult != tt.wantResult {
			t.Errorf("CalculateMulitplyTwo(%v) = %v, want %v", tt.input, gotResult, tt.wantResult)
		}
	}
}
func TestMultiply(t *testing.T) {
	if CalculateDivisionTwo(0) != 0 {
		t.Error("PackageTest: Failed")
	}
}
