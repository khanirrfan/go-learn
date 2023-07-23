package main

import (
	"fmt"
	"time"
)

func displayString(str string){
	for i := 0; i<5; i++{
		time.Sleep(1 * time.Second);
		fmt.Println(str)
	}
}
/* ad
	. goroutine are cheaper than threads
	. goroutine stored in stack and can grow dynamically
	  as compared to threads
	. to prevent race condition goroutines are best
	  
*/

func main(){

 go displayString("Welcome")

 //  goroutne on anonymous funtion
// go func (str string){
// 	for i := 0; i<5; i++{
// 		time.Sleep(1 * time.Second);
// 		fmt.Println(str)
// 	}
// }

 displayString("Lets learn")
}