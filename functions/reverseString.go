package main

import (
	"fmt"
)

func main() {
	str := "yoseaph"

	strRev := reverseString(str)
	fmt.Println(strRev)

}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}

	return
}
