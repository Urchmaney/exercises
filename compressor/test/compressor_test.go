package compressor_test

import (
	"fmt"
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

func TestGenerateHuffmanBinaryTreeFromFrequency(t *testing.T) {
	t.Run("Minimal", func(t *testing.T) {
		data := map[rune]int{'C': 32, 'D': 42, 'E': 120, 'K': 7, 'L': 42, 'M': 24, 'U': 37, 'Z': 2}
		expected := 306

		result := compressor.GenerateHuffmanBinaryTreeFromFrequency(data)
		fmt.Println(result)
		if *result.val > expected {
			t.Errorf("Expected %v but got %v", expected, 2)
		}
	})
}
