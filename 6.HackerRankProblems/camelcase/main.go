package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'camelcase' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */
func camelcase(s string) int32 {
	counter := 1 // Returns number of words - starts at 1 necessarily.
	r := []rune(s)

	// Runes for capital letters are between 65 and 90 (A-Z). Any rune in this range should increment the word counter since it indicates the start of a new word.
	for _, char := range r {
		if char > 64 && char < 91 {
			counter++
		}
	}

	return int32(counter)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	str, bl := os.LookupEnv("OUTPUT_PATH")
	if bl {
		fmt.Println(str)
	}

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := camelcase(s)

	fmt.Fprintf(writer, "%d\n", result)

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
