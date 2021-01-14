package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
	1 3 4 2 5 / 1 3 4 2 5 / 1 2 4 3 5 / 1 2 3 4 5
	1 3 5 2 4 6 7/1 2 5 3
	4 3 1 2 / 1 3 4 2 / 1 3 2 4 / 1 2 3 4
*/
func min(arr []int32) int32 {
	min := int32(math.MaxInt32)
	idxMin := 0
	for idx, e := range arr {
		if e < min {
			min = e
			idxMin = idx
		}
	}
	return int32(idxMin)
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	if len(arr) < 2 {
		return 0
	}
	idxMin := min(arr[1:])
	counter := 0
	if arr[0] > arr[idxMin+1] {
		temp := arr[0]
		arr[0] = arr[idxMin+1]
		arr[idxMin+1] = temp
		counter++
	}
	// fmt.Printf("sorted is %v counter is %d\n", arr, counter)
	return minimumSwaps(arr[1:]) + int32(counter)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)
	// fmt.Println(res)
	fmt.Fprintf(writer, "%d\n", res)

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
