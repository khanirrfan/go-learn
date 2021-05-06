package main

import "fmt"

func main(){
	// closure is the combination of a function and its sorruondings.
	newInts := getSum();
	 fmt.Println(newInts())
	 fmt.Println(newInts())
	 fmt.Println(newInts())
	 fmt.Println(newInts())
	 fmt.Println(newInts())
	 
	anotherNewInt := getSum()
	fmt.Println(anotherNewInt())
}
// garbage collector is method which removes the unused values or used.
func getSum() func() int{
	i := 0
   return func() int {
	   i++
	   return i;
   }

}