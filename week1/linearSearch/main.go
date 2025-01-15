package main

import "fmt"

func search(l []int, d int) (int, bool) {
	for i, v := range l {
		if d == v {
			return i, true
		}
		fmt.Printf("Checking %v at position %v\n", v, i)
		fmt.Println("--------------------------------")

	}
	return -1, false
}

func main() {
	numS := []int{1, 2, 3, 4}
	num := 44
	element, cond := search(numS, num)
	if cond == false {
		fmt.Printf("Couldn't find %v\n", num)
		return
	}
	fmt.Printf("Found %v at position %v\n", num, element)
}
