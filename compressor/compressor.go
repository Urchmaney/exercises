package compressor

import (
	"fmt"
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
