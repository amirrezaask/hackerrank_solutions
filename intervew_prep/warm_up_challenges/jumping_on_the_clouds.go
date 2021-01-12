package main

import "fmt"

func findPath(clouds []int32, currentCloud int32, path []int32) []int32 {
	if len(clouds) == 1 {
		return append(path, currentCloud)
	}
	var tempPath1 []int32
	var tempPath2 []int32

	if len(clouds) > 1 && clouds[1] == 0 {
		tempPath1 = findPath(clouds[1:], currentCloud+1, append(path, currentCloud))
	}
	if len(clouds) > 2 && clouds[2] == 0 {
		tempPath2 = findPath(clouds[2:], currentCloud+2, append(path, currentCloud))
	}
	if tempPath1 == nil {
		return tempPath2
	}
	if tempPath2 == nil {
		return tempPath1
	}
	if len(tempPath1) < len(tempPath2) {
		return tempPath1
	}
	return tempPath2
}

func main() {
	// 0,0,0,0,0
	l := []int32{0, 0, 1, 0, 0, 1, 0}
	fmt.Println(findPath(l, 0, []int32{}))
}
