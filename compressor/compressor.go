package compressor

import (
	"fmt"
	"sort"
)

type Node struct {
	val   int
	char  rune
	left  *Node
	right *Node
}

func FrequencyCalculator(val string, freq *map[rune]int) {
	for _, s := range val {
		(*freq)[s] += 1
	}
}

func GenerateHuffmanBinaryTreeFromFrequency(freq map[rune]int) Node {
	var nodeArr []Node
	for k, v := range freq {
		nodeArr = append(nodeArr, Node{val: v, char: k})
	}
	sort.Slice(nodeArr, func(i, j int) bool {
		return nodeArr[i].val < nodeArr[j].val
	})

	for len(nodeArr) > 1 {
		newNode := Node{val: nodeArr[0].val + nodeArr[1].val, left: &nodeArr[0], right: &nodeArr[1]}

		nodeArr = addNodeToSortedPosition(newNode, nodeArr[2:])
	}
	for _, v := range nodeArr {
		fmt.Println(v.val)
	}
	return nodeArr[0]
}

func addNodeToSortedPosition(n Node, arr []Node) []Node {
	var position int
	for ; position < len(arr); position++ {
		if arr[position].val > n.val {
			break
		}
	}
	result := append([]Node{}, arr[:position]...)
	result = append(result, n)
	return append(result, arr[position:]...)
}
