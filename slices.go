package main

import "fmt"

func main(){
	s := make([]string, 3) //syntax for slices
	s[0] = "100"
	s[1] = "200"
	s[2] = "300"

	fmt.Println(s)
	fmt.Println(s[2])
	// append value to slices

	s = append(s, "400") // a.push("400") 
	fmt.Println(s)
	// // copy array to some other array
	copyArray := make([]string, len(s))
	copy(copyArray, s)
	fmt.Println(copyArray);

	l := s[2:]
	fmt.Println(l)
	beforeL := s[:2]
	fmt.Println(beforeL);

	// two dimesional slices

	twoDSlices := make([][]int, 4)
}
