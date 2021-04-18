package main

import "fmt"

func main() {
	arr := []int{1,3,4,7,8} 
	fmt.Println(arr);
	fmt.Println(arr[0]);
	for i := 0; i<len(arr); i++ { // i++ => i = i+1 ===> i = 4 + 1 = 5
		if i%2 == 0 {
			fmt.Println(i);
		} else {
			fmt.Println("Number is not even")
		}
	}
}