package main

import "fmt"

type details struct{
	length int;
	width int
}

type myInterface interface{
	area() int
	perimeter() int
}

func (p details) area() int {
fmt.Println("Hello from area")
return p.length * p.width
}


func (p details) perimeter() int{
	return 2 * (p.length + p.width)
}

func main(){

	var t myInterface
	t = details{20,30,}
	fmt.Println(t.area())
	fmt.Println(t.perimeter())
}