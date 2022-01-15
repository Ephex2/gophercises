package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	// Capital letters exist in the range of 65-90 when the runes are converted to int
	// Lower-case letters exist in the range of 97-122 when the runes are converted to int32
	var caesarText string
	r := []rune(s)
	k = k % 26 // shifting 27 to the right is the same as shifting 1, apply this conversion here

	for _, char := range r {
		if char > 64 && char < 91 {
			shiftedChar := char + k
			if shiftedChar > 90 { // 90 is the upper-limit value for lowercase letters
				shiftedChar = shiftedChar - 26
			}

			caesarText += string(shiftedChar)
		} else if char > 96 && char < 123 {
			shiftedChar := char + k
			if shiftedChar > 122 { // 122 is the upper-limit value for uppercase letters
				shiftedChar = shiftedChar - 26
			}

			caesarText += string(shiftedChar)
		} else {
			caesarText += string(char)
		}
	}

	return caesarText
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	if n < 0 {
		fmt.Println("Needed to do something with n to force compile. For some reason ints are being entered as first input. let's see if this is enough.")
	}

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

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
