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
	ch := make(chan string)
	 go anotherFunc(ch)
	 res, ok := <-ch;
	 fmt.Println(res);
	 if ok == false {
		fmt.Println("closed")
	 }
	for i := 0; i < 4; i++ {
		cha <- " Irfan "
	}
	close(cha)
	fmt.Println("second channel closed");
}

func anotherFunc(cha chan string){
	close(cha);
}
