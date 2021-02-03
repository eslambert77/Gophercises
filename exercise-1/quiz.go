package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func stdErr(err error) {
	if err != nil {
		fmt.Println("Error!")
	}
}

func main() {
	//create a new csvReader
	pf, err := os.Open("problems.csv")
	stdErr(err)
	defer pf.Close()

	reader := csv.NewReader(pf)
	questions, err := reader.ReadAll()
	stdErr(err)

	score := 0
	input := bufio.NewReader(os.Stdin)
	for _, question := range questions {
		fmt.Print(question[0] + ": ")
		a, _, err := input.ReadLine()
		stdErr(err)
		if question[1] == string(a) {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d", score, len(questions))
}
