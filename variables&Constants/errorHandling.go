package main

import (
	"errors"
	"fmt"
)


func main(){
	params := []int{12,13,45}

for _, v := range params{
	r, e := funcError(v)
	if e != nil {
		fmt.Println("function failed to give output", e)
	} else {
		fmt.Println("function is working smoothly", r)
	}
}

}


func funcError(args int)(int, error){
	if args == 13 {
		return -1, errors.New("Can not process ahead with args value 13")
	}
	return args, nil
}