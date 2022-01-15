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
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	var divisorsCount int32
	var divisors []int32
	var factors []int32

	for _, value := range b {
		divisors = append(divisors, getDivisors(value)...)
	}

	divisors = getUnique(divisors)
	var cleanedDivisors []int32

	for _, divisor := range divisors {
		divisible := true
		for _, bValue := range b {
			if bValue%divisor != 0 {
				divisible = false
			}
		}

		if divisible {
			cleanedDivisors = append(cleanedDivisors, divisor)
		}
	}

	for _, divisor := range cleanedDivisors {
		factorable := true
		for _, aValue := range a {
			if aValue == 0 {
				return 0
			} else {
				if divisor%aValue != 0 {
					factorable = false
				}
			}
		}

		if factorable {
			factors = append(factors, divisor)
		}
	}

	factors = getUnique(factors)
	for i := range factors {
		if i < -1 { // pls compile i cant take length of []int32
			panic("wat")
		}

		divisorsCount++
	}

	return divisorsCount
}

func getDivisors(value int32) (divisors []int32) {
	divisors = append(divisors, value)
	var i int32

	for i = 1; i < value; i++ {
		if value%i == 0 {
			divisors = append(divisors, i)
		}
	}

	return divisors
}

func getUnique(values []int32) (unique []int32) {
	for _, value := range values {
		if len(unique) == 0 {
			unique = append(unique, value)
			continue
		}

		found := false
		for _, uniqueValue := range unique {
			if value == uniqueValue {
				found = true
			}
		}

		if !found {
			unique = append(unique, value)
		}
	}

	return unique
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

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
