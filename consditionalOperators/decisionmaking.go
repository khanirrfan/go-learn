// package main
//
// import "fmt"
//
// func main(){
// 	marks := 30
// 	if marks >= 30 {
// 		university := "IIIT"
// 		if university == "IIT" {
// 			fmt.Println(university)
// 		} else {
// 			fmt.Println("university is different")
// 		}
// 		fmt.Println("Student is selected")
// 	} else {
// 		fmt.Println("Student is not selected")
// 	}
// }
//
package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var da, ma, ya int64 = 9, 6, 20 // expected return data
	var de, me, ye int64 = 6, 6, 20 //due date
	var fine, late int64 
	// fmt.Scan(&da)
	// fmt.Scan(&ma)
	// fmt.Scan(&ya)
	// fmt.Scan(&de)
	// fmt.Scan(&me)
	// fmt.Scan(&ye)
	if ya <= ye { 
		if ma <= me {
			if da <= de { 
				fine = 0
				fmt.Println(fine)
			} else if da >= de {
				late = da - de
				fine = 15 * late
				fmt.Println(fine)
			} else {
				fmt.Println(("there is no fine"))
			}
		} else if ma > me {
			if ya < ye {
				late = ma - me
				fine = 0
				fmt.Println(fine)
			} else if ya == ye {
				late = ma - me
				fine = 500 * late
				fmt.Println(fine)
			}
		}
	} else {
		fine = 10000
		fmt.Println(fine)
	}
}