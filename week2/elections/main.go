package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	J = "James"
	S = "Susan"
)
var chances = 3

func vote(q string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	j := 0
	s := 0
	fmt.Println(q)
	vote, err := reader.ReadString('\n')
	vote = strings.Trim(vote, "\n")
	if err != nil || !strings.EqualFold(vote, J) || !strings.EqualFold(vote, S) {

		log.Printf("error: %v", err.Error())
		return 0, err
	}

	if strings.EqualFold(vote, J) {
		j++
		return 2, nil
		// os.WriteFile("votes.txt", )
	} else if strings.EqualFold(vote, S) {
		s++
		return 1, nil
	}
	chancesLeft(chances)
	return 0, nil
}

func main() {
	log.Printf("Chances allowed: %v\n", chances)
	quiz := "Your vote?: "
	if chances > 0 {
		fmt.Println(vote(quiz))
	}
}

func chancesLeft(chances int) {
	chances--
	log.Printf("Chances allowed: %v\n", chances)
}
