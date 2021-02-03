package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func stdErr(err error) {
	if err != nil {
		fmt.Println("Error!")
	}
}

func handleScore(s, quit <-chan int) int {
	score := 0
	for {
		select {
		case <-s:
			score++
		case <-time.After(time.Duration(duration) * time.Second):
			return score
		case <-quit:
			return score
		}
	}
}

var duration int
var filePath string

func init() {
	flag.IntVar(&duration, "timer", 30, "Sets the timer limit per question (default: 30)")
	flag.StringVar(&filePath, "file", "problems.csv", "Sets the file to open for questions(default: problems.csv)")
}

func main() {
	flag.Parse()
	//create a new csvReader
	pf, err := os.Open(filePath)
	stdErr(err)
	defer pf.Close()

	reader := csv.NewReader(pf)
	questions, err := reader.ReadAll()
	stdErr(err)

	score := make(chan int)
	quit := make(chan int)
	input := bufio.NewReader(os.Stdin)
	go func() {
		for _, question := range questions {
			fmt.Print(question[0] + ": ")
			a, _, err := input.ReadLine()
			stdErr(err)
			if strings.TrimSpace(question[1]) == string(a) {
				score <- 1
			}
		}
		quit <- 0
	}()

	finalScore := handleScore(score, quit)
	fmt.Printf("You scored %d out of %d questions.", finalScore, len(questions))
}
