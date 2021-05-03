package main

import "fmt"


func main(){
	// range iterates over elements ina variety of data structure
// index		 0 1 2 3
	num := []int{2,5,7,3}
	sum := 0
	// len(num) = 4
	// 2<4 - true
	// 3<4 - true
	// 4<4 - false
	// for i:=0; i<len(num); i++{
		// sum = sum + num[i] // 0 + num[0] = 0+2 = 2
						  //  2 + num[1] = 2 + 5 = 7
						//   7 + num[2] = 7 + 7 = 14
						// 14 + num[3] = 14 +3 = 17
		// sum = 17
		// i++ // i= i + 1 = 3 + 1 = 4
	// }

	for _, value := range num{
		// fmt.Println(index, value)
		sum = sum + value
	}
	fmt.Println(sum)

	// range applies to maps as well

	kvp := map[string]string{"a":"apple", "b": "banana"}
	for k, v := range kvp {
		fmt.Printf("%s -.> %s \n", k, v)
	}
}