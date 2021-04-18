// operators

package main

import "fmt"

func main(){

	// there many types of operators
	a := 20
	b := 30
	// 1. Arithmetic operators( +, -, /, *, %)
	
	fmt.Println(a +b ) // result 50
	fmt.Println( a - b) 

	// 2. Relational Operator ( ==, !=, >, <)
	fmt.Println(a == b)
	fmt.Println( a != b)
	fmt.Println( a > b)
	fmt.Println( a < b)

	// 3. Logical operators (&&, ||, !)
	//  a && b = true
// 	if !a{
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

	a++ // a = a+1
	b--

	a += 1 // a = a+1
	a *= 2 // a = a*2
}