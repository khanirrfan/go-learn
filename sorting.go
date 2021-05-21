package main

import "fmt"

func main() {
	a := []int{4, 5, 3, 2, 1, 7}
	n := len(a) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = swap(a[j], a[j+1])
			}
		}
	}

	fmt.Println(a)
}

func swap(a int, b int) (x int, y int) {
	x = b
	y = a
	return x, y
}
