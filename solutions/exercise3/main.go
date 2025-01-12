package main

import (
	"fmt"
)

const Five = 5

func lessThanFive(numS []int) []int {
	newNums := make([]int, 0, 5)
	for _, v := range numS {
		if v <= Five {
			newNums = append(newNums, v)
		}
	}
	return newNums
}

func main() {
	numS := []int{100, 1, 5, 60, 4, 20, 3, -9}
	digits := lessThanFive(numS)
	fmt.Printf("Numbers less than 5 in the array: %v\n", digits)
}
