package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	num, err := AskForNum()
	if err != nil {
		log.Fatal(err)
	}
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func AskForNum() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input a random positive number: ")
	num, err := reader.ReadString('\n')
	if err != nil {
		return 0, errors.New(err.Error())
	}
	num = strings.TrimSpace(num)
	newNum, err2 := strconv.Atoi(num)
	if err2 != nil || newNum <= 0 || newNum > 1000000000000000 {
		return 0, errors.New("wrong number format")
	}
	return newNum, nil
}
