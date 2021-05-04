package main

import "fmt"

func main(){
	i := 1
fmt.Println("initial:", i)
zeroValue(i);
fmt.Println("zeroValue:", i)
zeroPtr(&i)
fmt.Println("zeroValue:", i)
fmt.Println("zeroPtr:", &i)
}


func zeroValue(val int){
	val = 0
}

func zeroPtr(val *int){
	*val = 0
}

