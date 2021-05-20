package main

import "fmt"

func main() {
	ch := make(chan string)

	go myFunction(ch)

	for {
		response, ok := <-ch
		if ok == false {
			fmt.Println("chanel is closed", ok)
			break
		}
		fmt.Println("channel is open", response)
	}
}

func myFunction(cha chan string) {
	//  go anotherFunc()
	for i := 0; i < 4; i++ {
		cha <- " Irfan "
	}
	close(cha)
}

// func anotherFunc(){
//
// }
