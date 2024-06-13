package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	compressor "../../../go-exercises/golang/compressor"
)

func main() {
	c := flag.String("f", "", "File Name")
	// o := flag.String("o", "", "Output File Name")
	flag.Parse()
	fileName := *c
	// outputFile := *o
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
	// compressor.AddPrefixCode(freq, root)
	// compressor.WriteFileHeader(freq, outputFile)
	// compressor.EncryptInputFile(fmt.Sprintf("../%v", fileName), outputFile, freq)
	compressor.DecryptCompressedFile("output.txt", root)
}
