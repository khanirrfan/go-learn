package main

import "fmt"

// function expression - sombody expression thought about something to third person
// /function statment
// anonymous function
// function declaration

func main(){
	// 
	result := sum(1,2,3)
	fmt.Println(result)

	result1, result2 := multiVals()

	fmt.Println(result1, result2);
	variadicFunc(1,3)
	variadicFunc(12,45,56)
	variadicFunc(12,45,56,34,5,6,1434)
}

func sum(a, b,c int) int{
	return a + b + c
}


// multiple return values

func multiVals() (int, int){
	return 3,7
}

// variadic functions = can be called with any numbe rof trailing arguments
// args = [a,b,c,d,e]
func variadicFunc(args ...int){

	fmt.Println(args);
	sum := 0
	for _, value := range args{
		sum = sum + value
	}

	fmt.Println(sum)
}
