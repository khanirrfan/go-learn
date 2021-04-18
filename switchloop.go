package main

import "fmt"

func main() {
	i := 1 // value of i is 2
	switch i { // i = 1
	case 1 :   // i == 1
	fmt.Println("Color code")
	case 2 :
		fmt.Println("Value found")
	default : 
	fmt.Println("Default color")	
	}
}