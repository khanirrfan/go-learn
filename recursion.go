// recursion- the function which calls itself until it reaches the final condition.

package main

import "fmt"


func main(){
	// 5 = 5*4*3*2*1*0= 120
	fmt.Println(fact(5))
}


func fact(n int) int{
	if n == 0 {
		return 1
	}
		// 5*4*3*2*1*0= 120
	//  5*fact(4)
			// 4*fact(3)
			// 	3*fact(2)
			// 		2*fact(1)
			// 			1*fact(0)
	return n*fact(n-1)
}