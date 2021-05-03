// Arrays

package main

import "fmt"

func main(){
	// one dimesional array
	// Array is a numbered sequence of elements of a specific length
	// here we have defined array with length of 5
	var arr [5]int

	// arr := []int{1,2,3,4}
	fmt.Println("array", arr)
	// here we have assigned value to 0th index of arr 
	arr[0] = 100
	fmt.Println("array", arr)
	// here we have defined array with values 
	var arr1 = []int {1,2,3,4,5}

	fmt.Println(arr1)
	// we have created empty array using sencond method of defining  arrays
	var arr2 = []int {}

	fmt.Println(arr2)

	// two dimensional arrays
	var twoDArray [2][3]int
	// twoDArray[0][0] = 1
	
	twoDArray[1][2] = 89
	// 	   columns

	// 		  |
	// row - [0 0 0 = 0
	//        1 1 1] = 1
	// let i var i
		for i := 0; i<2; i++ { 
			for j := 0; j<3; j++ {
				twoDArray[i][j] = i +j
			}
		}
	fmt.Println(twoDArray)

}