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

// credits for this algorigthm goes to https://www.hackerrank.com/amansbhandari
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int64(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int64(mTemp)

	arr := make([]int64, n+1)
	var x int64
	max := int64(math.MinInt64)
	for i := int64(0); i < m; i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int64
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesRow = append(queriesRow, queriesItemTemp)

		}
		if len(queriesRow) != int(3) {
			panic("Bad input")
		}
		p := queriesRow[0]
		q := queriesRow[1]
		sum := queriesRow[2]

		arr[p] += sum
		if q+1 <= n {
			arr[q+1] -= sum
		}
	}
	for i := int64(1); i <= n; i++ {
		x += arr[i]
		if x > max {
			max = x
		}
	}
	// fmt.Println(max)
	fmt.Fprintf(writer, "%d\n", max)

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
