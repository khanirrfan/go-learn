package main

import "fmt"

// syntax of methods

// func receiver() functionName() return type {}
type person struct{
	firstName string
	lastName string
	age int
}

type data int
// non-struct type recevier
func (d1 data) getPersonDetail(){

	fmt.Println(d1)

}

// struct type recevier
// func (p person) getPersonDetail(){
// 
// 	fmt.Println(p.firstname)
// 
// }

// pointer struct type recevier
// func (p *person) getPersonDetail(){
// 
// 	fmt.Println(p.firstname)
// 
// }
func main() {
	// p := person{
	// 	"irfan",
	// 	"khan",
	// 	27,
	// }
	// res := &p
	value1 := data(23)
	value1.getPersonDetail()
	// res.getPersonDetail()
	// p.getPersonDetail()
}