package main

import "fmt"

func main(){
	s := make([]string, 3)
	s[0] = "100"
	s[1] = "200"
	s[2] = "300"

	fmt.Println(s)
	// fmt.Println(s[2])
	// append value to slices

	s = append(s, "400") // a.push("400") 
	fmt.Println(s)
	// copy array to some other array
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c);
}