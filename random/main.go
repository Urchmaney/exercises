package main

import (
	"fmt"
)

func main() {
	val := recursive_sum([]int {1,2,3,4,5,6,7,8,9, 0})
	fmt.Println(val)
}


func sum(values ...int) int {
	total := 0

	for _, v := range values {
		total += v
	}
	return total
}

func recursive_sum(values []int) int {
	if len(values) <= 0 { return 0 }

	return values[0] + recursive_sum(values[1:])
}

