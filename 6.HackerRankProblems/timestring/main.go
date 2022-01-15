package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	timeChar := s[len(s)-2]
	cleanedString := s[0 : len(s)-2]
	var outputString string

	timeArr := strings.Split(cleanedString, ":")

	if timeChar == 'P' {
		intHour, err := strconv.Atoi(timeArr[0])
		if err != nil {
			log.Fatal(err.Error())
		}

		hourValue := (intHour + 12) % 24
		timeArr[0] = fmt.Sprint(hourValue)
	}

	for i, timeValue := range timeArr {
		var stringBit string

		if i != 0 {
			stringBit += ":"
		}

		stringBit += timeValue
		outputString += stringBit
	}

	return outputString
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
