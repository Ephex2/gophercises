package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

var birdTypeDensityMap = make(map[int]int)

func migratoryBirds(arr []int) int32 {
	if len(arr) == 1 {
		return int32(arr[0])
	}

	sort.Ints(arr)
	var hitCounter = 1

	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			hitCounter++
		} else {
			birdTypeDensityMap[arr[i-1]] = hitCounter
			hitCounter = 1
		}

		if i == len(arr)-1 {
			birdTypeDensityMap[arr[i]] = hitCounter
		}
	}

	var highestMode int
	var mostCommonSmallType int32
	fmt.Println(birdTypeDensityMap)
	for i, j := range birdTypeDensityMap {
		if j > highestMode {
			fmt.Printf("Setting mostCommonSmallType to: %v in block 1. highestMode is: %v, j is: %v\n", int32(i), highestMode, j)
			highestMode = j
			mostCommonSmallType = int32(i)
		} else if j == highestMode && int32(i) < mostCommonSmallType {
			fmt.Printf("Setting mostCommonSmallType to: %v in block 2. highestMode is: %v, j is: %v\n", int32(i), highestMode, j)
			mostCommonSmallType = int32(i)
		}
	}

	return mostCommonSmallType
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := migratoryBirds(arr)
	fmt.Println(result)
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
