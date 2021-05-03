// data types

package main

import (
	"fmt"
) 

func main(){
	fmt.Println("samlkdas");

// every line in a program is an expression.
// every expression is made of tokens, consists of keyword, identifiers, string constant, sysmbol etc

	// var arrayNumbers int32 = 234
	// data types are 4 types

	// 1. Boolean type = true/false

	// 2. Numeric types = int/float
		// var - variable name will not start form any special character and number, it will start from A-Z, a-z, _
		
		// valid declaration
		// var Another int8
		// var another uint8
		// var _another unit 16

		// invalid declaration
		// var 1another, @another - can not declare variable like this
		// uint, int
			// uint8, uint16, uint32, uint64

					// uint - are only +ve number starting from 0 to ....
					// int - are both -ve and +ve numbers - t0 +
						// uint8 - 2^8 = 0 to 255; // var nameUint uint8 = 1024
						// uint16 - 2^16 = 0 to 
						// uint32 - 2^32 = 0 to
						// uint64 - 2^64 = 0 to 

						// int8 - 2^8 = -128 to +128
						// int16 - 2^16 = - to +
						// int32 - 2^32 = - to +
						// int64 - 2^64 = - to +

		// float8, float16, float32, float64

	// 3. String type
		var anotherString string = "this is string"; 
		secodnString := "this is second string declaration" //- short hand methoid/property

	// derived types - array, linked list, json, struct, function, slices types

	// type objectName struct {
	// 	firstname : 'irfan',
	// 	lastname:'khan'
	// }

	// declaration

		// var arr []String, []int32, []uint32 = [1,2,3,4,]
		// function types - 

		// func anotherFunction() []int16{
		// 	return []int16
		// }
		}