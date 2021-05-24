package main

import "fmt"

func main() {
	for j := 2; j <= 10; j++ {
		for i := 1; i <= 10; i++ {
			fmt.Println(j, "*", i, "=", j*i)
		}
	}
}
