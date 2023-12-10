package stress

import (
	"math"
	"testing"
)

func TestMeanStress(t *testing.T) {
	tests := []struct {
		minValue, maxValue float64
		expected           float64
		expectError        bool
	}{
		{1, 5, 3, false},
		{5, 1, 0, true},
		{10, 10, 10, false},
		{-5, 5, 0, false},
		{10, 20, 15, false},
		{-45, -20, 0, false},
	}

	for _, tt := range tests {
		got, err := MeanStress(tt.minValue, tt.maxValue)
		if tt.expectError {
			if err == nil {
				t.Errorf("Mean stress(%.4f, %.4f) expected an error, but got none", tt.minValue, tt.maxValue)
			}
			continue
		}
		if err != nil {
			t.Errorf("Mean stress(%.4f, %.4f) returned an unexpected error: %.4f", tt.minValue, tt.maxValue, err)
		}
		if !almostEqual(got, tt.expected) {
			t.Errorf("Mean stress(%.4f, %.4f) = %.4f, want %.4f", tt.minValue, tt.maxValue, got, tt.expected)
		}
	}
}

func TestAlternatingStress(t *testing.T) {
	tests := []struct {
		minValue, maxValue float64
		expected           float64
		expectError        bool
	}{
		{-20, 20, 20, false},
		{5, -10, 0, true},
		{0, 10, 5, false},
		{-5, 5, 5, false},
		{-34.34, 45.53, 39.935, false},
	}

	for _, tt := range tests {
		got, err := AlternatingStress(tt.minValue, tt.maxValue)
		if tt.expectError {
			if err == nil {
				t.Errorf("Alternating stress(%.4f, %.4f) expected an error, but got none", tt.minValue, tt.maxValue)
			}
			continue
		}
		if err != nil {
			t.Errorf("Alternating stress(%.4f, %.4f) returned an unexpected error: %.4f", tt.minValue, tt.maxValue, err)
		}
		if !almostEqual(got, tt.expected) {
			t.Errorf("Alternating stress(%.4f, %.4f) = %.4f, want %.4f", tt.minValue, tt.maxValue, got, tt.expected)
		}
	}
}

func TestGoodmanStress(t *testing.T) {
	tests := []struct {
		meanStress, alternatingStress, ultimateStrength float64
		expected                                        float64
		expectError                                     bool
	}{
		{-20, 20, -10, 20, true},
		{40, 20, 120, 30, false},
		{10, 45.23, 100, 50.2556, false},
		{23.56, 65.6, 120, 81.6259, false},
	}

	for _, tt := range tests {
		got, err := GoodmanStress(tt.meanStress, tt.alternatingStress, tt.ultimateStrength)
		if tt.expectError {
			if err == nil {
				t.Errorf("Goodman stress(%.4f, %.4f, %.4f) expected an error, but got none", tt.meanStress, tt.alternatingStress, tt.ultimateStrength)
			}
			continue
		}
		if err != nil {
			t.Errorf("Goodman stress(%.4f, %.4f, %.4f) returned an unexpected error: %.4f", tt.meanStress, tt.alternatingStress, tt.ultimateStrength, err)
		}
		if !almostEqual(got, tt.expected) {
			t.Errorf("Goodman stress(%.4f, %.4f, %.4f) = %.4f, want %.4f", tt.meanStress, tt.alternatingStress, tt.ultimateStrength, got, tt.expected)
		}
	}
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-4
}
