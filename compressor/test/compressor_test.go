package compressor_test

import (
	"testing"

	compressor "../../../go-exercises/compressor"
)

func TestFrequencyCounter(t *testing.T) {

	t.Run("Calculate Frequency", func(t *testing.T) {
		data := "aaabbbc"
		result := make(map[rune]int)

		if len(result) != 0 {
			t.Errorf("Expected 0 but got %v", len(result))
		}

		compressor.FrequencyCalculator(data, &result)
		expected := 3

		if len(result) != expected {
			t.Errorf("Expected %v but got %v", expected, len(result))
		}
	})

}
