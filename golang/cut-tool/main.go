package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	cut "./cut"
)

func main() {
	f := flag.String("f", "", "Field Index")
	d := flag.String("d", "\t", "Delimeter")

	flag.Parse()
	fileName := flag.Arg(0)

	fieldStr := *f
	var fields []string
	if strings.Contains(fieldStr, ",") {
		fields = strings.Split(fieldStr, ",")
	} else {
		fields = strings.Split(fieldStr, " ")
	}
	delimeter := *d

	var file *os.File
	if isPiped() {
		file = os.Stdin
	} else {
		ffile, err := os.Open(fileName)
		if err != nil {
			panic(fmt.Sprintln("Error Opening file", err))
		}
		file = ffile
	}

	buffer := make([]byte, 32*1024)

	for {
		_, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading file from file", err)
			return
		}
		splitted := bytes.Split(buffer, []byte{'\n'})
		for _, v := range splitted {
			vSplit := bytes.Split(v, []byte(delimeter))
			fieldsResult := make([]string, len(fields))
			for i, f := range fields {
				index, err := strconv.ParseInt(f, 10, 0)
				if err != nil {
					panic("Error converting parameter")
				}

				fieldsResult[i] = string(vSplit[index-1])
			}
			fmt.Println(strings.Join(fieldsResult, delimeter))
		}
	}
	cut.Cut()
}

func isPiped() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return !(stat.Mode()&os.ModeNamedPipe == 0)
}
