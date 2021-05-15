package main


type Person struct{
	name string
	age int
}


type People struct{
	Person
	work string
}


// type embedding is inheriting or extending something to somewhere.
