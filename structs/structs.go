// structs are typed collection, they are like json objects in golang

package main

import "fmt"

type person struct{
	firstName string
	lastName string
	age int
}

func main(){
personDetails := person{
	firstName: "Irfan",
	lastName:"Khan",
	age: 27,
}

fmt.Println(personDetails)

	fmt.Println(newPerson("Yoseaph"))
}

func newPerson(name string) *person{
	p := person{firstName:name}
	p.lastName = "something"
	p.age = 18
	return &p
}