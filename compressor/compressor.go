package compressor

func FrequencyCalculator(val string, freq *map[rune]int) {
	for _, s := range val {
		(*freq)[s] += 1
	}
}
