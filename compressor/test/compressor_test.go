package compressor_test

import (
	"testing"

	compressor "../../../go-exercises/compressor"
)

func TestFrequencyCounter(t *testing.T) {

	t.Run("Calculate Frequency", func(t *testing.T) {
		data := "aaabbbc"
		result := make(map[rune]*compressor.CharFreq)

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
		data := map[rune]*compressor.CharFreq{
			'C': &compressor.CharFreq{Freq: 32, Code: ""},
			'D': &compressor.CharFreq{Freq: 42, Code: ""},
			'E': &compressor.CharFreq{Freq: 120, Code: ""},
			'K': &compressor.CharFreq{Freq: 7, Code: ""},
			'L': &compressor.CharFreq{Freq: 42, Code: ""},
			'M': &compressor.CharFreq{Freq: 24, Code: ""},
			'U': &compressor.CharFreq{Freq: 37, Code: ""},
			'Z': &compressor.CharFreq{Freq: 2, Code: ""}}
		expected := 306

		result := compressor.GenerateHuffmanBinaryTreeFromFrequency(data)
		if result.Val != expected {
			t.Errorf("Expected %v but got %v", expected, result.Val)
		}

		compressor.AddPrefixCode(data, result)
		for _, v := range data {
			if len(v.Code) == 0 {
				t.Errorf("Expected length to be greather than 1 but got %v", v.Code)
			}
			// fmt.Printf("%v     %v\n", string(k), v.Code)
		}

	})
}
