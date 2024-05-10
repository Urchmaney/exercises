package compressor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

type Node struct {
	Val   int
	Char  rune
	Left  *Node
	Right *Node
}

type CharFreq struct {
	Freq int
	Code string
}

func FrequencyCalculator(val string, freq *map[rune]*CharFreq) {
	for _, s := range val {
		data := (*freq)[s]
		if data == nil {
			(*freq)[s] = &CharFreq{Freq: 1, Code: ""}
			continue
		}
		data.Freq += 1
	}
}

func GenerateHuffmanBinaryTreeFromFrequency(freq map[rune]*CharFreq) Node {
	var nodeArr []Node
	for k, v := range freq {
		nodeArr = append(nodeArr, Node{Val: v.Freq, Char: k})
	}

	sort.Slice(nodeArr, func(i, j int) bool {
		return nodeArr[i].Val < nodeArr[j].Val
	})

	for len(nodeArr) > 1 {
		newNode := Node{Val: nodeArr[0].Val + nodeArr[1].Val, Left: &nodeArr[0], Right: &nodeArr[1]}

		nodeArr = addNodeToSortedPosition(newNode, nodeArr[2:])
	}

	if len(nodeArr) == 0 {
		return Node{Val: 0}
	}

	return nodeArr[0]
}

func AddPrefixCode(freq map[rune]*CharFreq, root Node) {
	for k, v := range freq {
		v.Code = *FindCode(&root, k)
	}
}

func FindCode(node *Node, char rune) *string {
	result := ""
	if node == nil {
		return nil
	}
	if node.Char == char {
		return &result
	}
	leftResult := FindCode(node.Left, char)
	rightResult := FindCode(node.Right, char)

	if leftResult == nil && rightResult == nil {
		return nil
	} else if leftResult != nil {
		result = fmt.Sprintf("%v%v", 0, *leftResult)
		return &result
	} else {
		result = fmt.Sprintf("%v%v", 1, *rightResult)
		return &result
	}
}

func WriteFileHeader(freq map[rune]*CharFreq, fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE, 0777)
	if err != nil {
		panic("Error Writing File Header")
	}
	defer file.Close()
	file.WriteString("=Header=>\n")
	for k, v := range freq {
		file.WriteString(fmt.Sprintf("%v|%v|%v\n", k, v.Freq, v.Code))
	}
	file.WriteString("<=Header=\n")
}

func EncryptInputFile(inputFile, outputFile string, lookup map[rune]*CharFreq) {
	fmt.Println("Encryption started")
	rfile, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Sprintf("Error Opening file with name '%v'", inputFile))
	}
	wfile, errw := os.OpenFile(outputFile, os.O_APPEND, 0777)
	if errw != nil {
		panic("Error Writing File Header")
	}
	defer wfile.Close()

	buffer := make([]byte, 32*1024)
	for {
		_, rerr := rfile.Read(buffer)
		for _, chr := range string(buffer) {
			code := lookup[chr].Code
			wfile.WriteString(code)
		}
		if rerr == io.EOF {
			break
		} else if rerr != nil {
			fmt.Println("Error encoding file from file", err)
			return
		}

	}

}

func DecryptCompressedFile(fileName string, node Node) {
	fmt.Println("Decryption Started=================")
	decryptFile := "decryted.txt"
	rfile, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error Opening file with name '%v'", fileName))
	}
	wfile, errw := os.OpenFile(decryptFile, os.O_CREATE, 0777)
	if errw != nil {
		panic("Error Writing File Header")
	}
	defer wfile.Close()

	buffer := make([]byte, 32*1024)
	hasFinihedHeader := false

	count := 0

	for {
		_, rerr := rfile.Read(buffer)
		if rerr == io.EOF {
			break
		} else if rerr != nil {
			fmt.Println("Error encoding file from file", rerr)
			return
		}
		if !hasFinihedHeader {
			splitted := bytes.Split(buffer, []byte{'\n'})

			for _, s := range splitted {
				data := string(s)
				if data == "<=Header=" {
					hasFinihedHeader = true
				}

				if hasFinihedHeader {
					decryptStringAndWrtite(*wfile, data, node)
				}
			}
		} else {
			decryptStringAndWrtite(*wfile, string(buffer), node)
		}

		count += 1
		fmt.Printf(" %v==\n", (count))

	}
	// _, rerr := rfile.Read(buffer)
	// if rerr != nil {
	// 	panic("Error")
	// }

	// splitted := bytes.Split(buffer, []byte{'\n'})

	// for _, s := range splitted {
	// 	data := string(s)
	// 	fmt.Println(data, "========")
	// 	if data == "<=Header=" {
	// 		fmt.Println("Foound")
	// 	}
	// }
}

func decryptStringAndWrtite(file os.File, s string, node Node) {
	sLen := len(s)
	index := 0
	for index < sLen {
		n := node
		for index < sLen {

			if s[index] == '0' {
				n = *n.Left
			} else {
				n = *n.Right
			}
			index += 1
			if n.Char != 0 {
				break
			}
		}
		file.WriteString(string(n.Char))
	}
	fmt.Println("end")
}

func addNodeToSortedPosition(n Node, arr []Node) []Node {
	var position int
	for ; position < len(arr); position++ {
		if arr[position].Val > n.Val {
			break
		}
	}
	result := append([]Node{}, arr[:position]...)
	result = append(result, n)
	return append(result, arr[position:]...)
}
