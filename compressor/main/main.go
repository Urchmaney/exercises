package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	compressor "../../../go-exercises/compressor"
)

func main() {
	c := flag.String("f", "", "File Name")
	flag.Parse()
	fileName := *c
	file, err := os.Open(fmt.Sprintf("../%v", fileName))
	if err != nil {
		panic(fmt.Sprintf("Error Opening file with name '%v'", fileName))
	}

	freq := make(map[rune]*compressor.CharFreq)

	buffer := make([]byte, 32*1024)

	for {
		_, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading file from file", err)
			return
		}

		compressor.FrequencyCalculator(string(buffer), &freq)
	}

	root := compressor.GenerateHuffmanBinaryTreeFromFrequency(freq)
	compressor.AddPrefixCode(freq, root)
	for _, v := range freq {
		fmt.Println(v.Code)
	}

}
