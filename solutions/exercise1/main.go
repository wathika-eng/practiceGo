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
	reader := bufio.NewReader(os.Stdin)
	name, age, err := AskNameAge(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %v, you have %d to be 100 years\n", name, 100-age)
}

func AskNameAge(reader *bufio.Reader) (string, int, error) {
	// create a reader using bufio
	// reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is your name?")
	name, err := reader.ReadString('\n')
	fmt.Printf("What is your age?")
	age, err2 := reader.ReadString('\n')
	if err != nil || err2 != nil {
		newErr := errors.Join(err, err2)
		return "", 0, errors.New(newErr.Error())
	}

	name = strings.TrimSpace(name)
	age = strings.TrimSpace(age)
	newAge, _ := strconv.Atoi(age)
	return name, newAge, nil
}
