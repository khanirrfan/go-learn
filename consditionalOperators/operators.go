// operators

package main

import "fmt"

func main(){

	// there many types of operators
	a := 20
	b := 30
	// 1. Arithmetic operators( +, -, /, *, %)
	
	fmt.Println(a + b ) // result 50
	fmt.Println( a - b) 

	// 2. Relational Operator ( ==, !=, >, <)
	// fmt.Println(a == b)
	fmt.Println( a != b) 
	// fmt.Println( a > b)
	// fmt.Println( a < b)

	// 3. Logical operators (&&, ||, !)
	//  a && b = true
	firstVar := "something"
	secondVar := "constant number"
	if firstVar == "2" && secondVar == "constant number" { // false && true = false
		fmt.Println("C is defned")
	} else {
		fmt.Println("c is not defined")
	}
// 	if !a{ !! !false = !true = false
// 
// 	}
// 	if a && b {
// 
// 	} else {
// 
// 	}
	// 4. Assignment operator (=, ++, --, *=,+=,-=)
	// a := 20 // value 20 is assigned to variable a using assignment operator( = )
	// b := 30

	//a++ // a = a+1 a = 20 + 1 = 21, 21+1 = 22, 22+1 = 23
	//b-- 

	//a += 1 // a = a+1
	//a *= 2 // a = a*2  = 20*2 = 40, 40*2 = 80
}