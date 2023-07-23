package main

import (
	"fmt"
	"strconv"
)

func main() {

	// for i := 0; i <= 10; i++ {
	// 	fmt.Print(strconv.Itoa(fiboSeries(i)) + " ")
	// }

	for i := 0; i <= 10; i++ {
		fmt.Print(strconv.Itoa(fiboLoopSeries(i)) + " ")
	}
}

func fiboSeries(num int) int {
	if num <= 1 {
		return num
	}
	return fiboSeries(num-1) + fiboSeries(num-2)
}

func fiboLoopSeries(num int) int {
	fSlice := make([]int, num+1, num+2)
	if num < 2 {
		fSlice = fSlice[0:2]
	}
	fSlice[0] = 0
	fSlice[1] = 1
	for i := 2; i <= num; i++ {
		fSlice[i] = fSlice[i-1] + fSlice[i-2] // 0 1 1 2
	}
	return fSlice[num]
}
