package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	inputPath = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit = flag.Int("limit", 30, "the time limit for the quiz in seconds")
)

func main() {
	flag.Parse()

	inputFile, _ := os.Open(*inputPath)
	defer inputFile.Close()

	csvReader := csv.NewReader(bufio.NewReader(inputFile))

	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	defer timer.Stop()

	questions, _ := csvReader.ReadAll()
	var result int


loop:
	for i, question := range questions {
		fmt.Printf("Problem #%d: %s = ", i+1, question[0])
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break loop
		case userAnswer := <-answerCh:
			if userAnswer == question[1] {
				result++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", result, len(questions))
}
