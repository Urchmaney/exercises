package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	c := flag.String("c", "", "File Name")
	l := flag.String("l", "", "File Name")
	w := flag.String("w", "", "File Name")
	m := flag.String("m", "", "File Name")

	flag.Parse()

	cFlag := *c
	lFlag := *l
	wFlag := *w
	mFlag := *m

	var fileName string
	isPipe := isPiped()
	var result bytes.Buffer

	if cFlag != "" {
		if !isPipe {
			fileName = cFlag
		}
		size := byteSize(cFlag)
		result.WriteString(fmt.Sprint(size))
		result.WriteString("  ")
	}

	if lFlag != "" {
		if !isPipe {
			fileName = lFlag
		}
		lineNumbers := fileLineNumber(lFlag, isPipe)
		result.WriteString(fmt.Sprint(lineNumbers))
		result.WriteString("  ")
	}

	if wFlag != "" {
		if !isPipe {
			fileName = wFlag
		}
		wordCount := wordCounts(wFlag, isPipe)
		result.WriteString(fmt.Sprint(wordCount))
		result.WriteString("  ")
	}

	if mFlag != "" {
		if !isPipe {
			fileName = mFlag
		}
		countCharacter := characterCounts(mFlag, isPipe)
		result.WriteString(fmt.Sprint(countCharacter))
		result.WriteString("  ")
	}

	if cFlag == "" && lFlag == "" && wFlag == "" && mFlag == "" {
		if !isPipe {
			fileName = "test.txt"
		}

		lineNumbers := fileLineNumber(fileName, isPipe)
		result.WriteString(fmt.Sprint(lineNumbers))
		result.WriteString("  ")

		wordCount := wordCounts(fileName, isPipe)
		result.WriteString(fmt.Sprint(wordCount))
		result.WriteString("  ")

		size := byteSize(fileName)
		result.WriteString(fmt.Sprint(size))
		result.WriteString("  ")

	}

	if fileName != "" {
		result.WriteString(fileName)
	}

	fmt.Println(result.String())

}

func isPiped() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return !(stat.Mode()&os.ModeNamedPipe == 0)
}

func getFile(fileName string, isPiped bool) (*os.File, error) {
	if isPiped {
		return os.Stdin, nil
	}

	return os.Open(fileName)
}

func byteSize(fileName string) int64 {
	file, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return file.Size()
}

func fileLineNumber(fileName string, isPiped bool) int {
	file, err := getFile(fileName, isPiped)
	if err != nil {
		fmt.Println("Error Opening your file", err)
	}
	buffer := make([]byte, 32*1024)
	count := 0
	lineSeperator := []byte{'\n'}

	for {
		c, err := file.Read(buffer)
		count += bytes.Count(buffer[:c], lineSeperator)
		if err == io.EOF {
			return count
		} else if err != nil {
			fmt.Println("Error reading file from file", err)
			return 0
		}
	}
}

func wordCounts(fileName string, isPiped bool) int {
	file, err := getFile(fileName, isPiped)
	if err != nil {
		fmt.Println("Error Opening your file", err)
	}
	buffer := make([]byte, 32*1024)
	count := 0
	wordSeperator := []byte{' '}
	for {
		c, err := file.Read(buffer)
		count += bytes.Count(buffer[:c], wordSeperator)
		if err == io.EOF {
			return count
		} else if err != nil {
			fmt.Println("Error reading file from file", err)
			return 0
		}
	}
}

func characterCounts(fileName string, isPiped bool) int {
	file, err := getFile(fileName, isPiped)
	if err != nil {
		fmt.Println("Error Opening your file", err)
	}
	buffer := make([]byte, 32*1024)
	count := 0
	characterSeperator := []byte("")

	for {
		c, err := file.Read(buffer)
		count += bytes.Count(buffer[:c], characterSeperator)
		if err == io.EOF {
			return count
		} else if err != nil {
			fmt.Println("Error reading file from file", err)
			return 0
		}
	}
}
