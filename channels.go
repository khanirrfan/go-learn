package main

import "fmt"

func main(){
// a channel is a medium in golang through a goroutine communicates with another goroutine and is commucation is lock-free.
// var channel_name type
myChannel := make(chan int)
// var myChannel chan int
// <- - sending information
// -> - recieving information
// fmt.Printf("%T", myChannel)
go funcOne(myChannel)
myChannel <- 100
}
func funcOne(myChan chan int){

	fmt.Println(200 + <-myChan)
	// syntax for closing channel
	
}

