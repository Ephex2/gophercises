package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var shuffle = false
var timerSeconds = 30
var reader = bufio.NewReader(os.Stdin)
var quizDone = make(chan struct{})
var totalCorrect = 0

func main() {
	// In a more serious project, consider bundling files with gobuffalo/packr or embedding files using embed
	filePath, err := filepath.Abs("C:/Git/gophercises/1.QuizGame/problems.csv")
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 2 {
		args := os.Args[1:]

		for i, arg := range args {
			arg = strings.ToLower(arg)

			switch arg {
			case "-shuffle":
				shuffle = true
			case "-filepath":
				f, err := filepath.Abs(args[i+1])
				if err != nil {
					panic(err)
				}

				filePath = f
			case "-timer":
				timerArg := args[i+1]
				t, err := strconv.Atoi(timerArg)
				if err != nil {
					fmt.Printf("The value: %v is not a valid number. Please enter a valid number to set the number of seconds on the timer\n", timerArg)
					return
				}

				timerSeconds = t
			default:
				// do nothing, don't want to write nonsense for the arguments that require values.
			}
		}
	}

	// Print arguments to aid debugging, consider deleting
	fmt.Printf("Shuffle: %v\n", shuffle)
	fmt.Printf("Quiz path: %v\n", filePath)
	fmt.Printf("TimerSeconds: %v\n", timerSeconds)

	problems := readCsvFile(filePath)
	total := len(problems)

	if shuffle {
		// Example code for shuffling a slice from: https://yourbasic.org/golang/shuffle-slice-array/
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
	}

	// setup timer
	fmt.Printf("The quiz and timer will begin after you press on the enter key - good luck!\n")
	reader.ReadString('\n')
	timer := time.NewTimer(time.Duration(timerSeconds) * time.Second)
	go QuizUser(problems, &totalCorrect)

	// it's a race against the clock; execution blocks here until one of the channels receives a message
	select {
	case <-timer.C:
		fmt.Printf("Time is up!\n")
		fmt.Printf("Scored %v out of %v questions.\n", totalCorrect, total)
	case <-quizDone:
		fmt.Printf("Scored %v out of %v questions.\n", totalCorrect, total)
	}

}

// Reads content from a CSV file into a slice of string slices. From:  https://stackoverflow.com/questions/24999079/reading-csv-file-in-go
// Changed log calls to panics since this exercise is simplistic.
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

// Separated from main to aid legibility and because I wanted to simply use a return in the select statement.
func QuizUser(problems [][]string, totalCorrect *int) {

	*totalCorrect = 0
	for i, problem := range problems {
		answer, err := strconv.Atoi(problem[1])
		if err != nil {
			fmt.Printf("Quiz answer for problem %v is not a number, skipping to next question.\n", i)
			continue
		}

		fmt.Printf("Problem %v: %v\n", i+1, problem[0])
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-2] // gets rid of newline and byte 13 character. Consider using trim or other scan method.
		userAnswer, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("Please enter a valid number. Answer rejected.\n")
		}

		if userAnswer == answer {
			*totalCorrect++
		}

	}

	quizDone <- struct{}{}
}
