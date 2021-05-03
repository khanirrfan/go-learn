package main

import "fmt"

func main(){
	//maps are go's built in associative data types. they store in values on key-pair pair.

	// array [ "", "", ""]
	//maps = {
	// 	name:"name1",
	// 	city:"city"
	// }
	// syntax for maps = make(map[string]string)
	m := make(map[string]int)
	m["key1"] = 21
	m["key2"] = 90
	fmt.Println("map:", m)
	// access invidual value by key names
	fmt.Println("key1 value:", m["key1"])
	// get the length of map
	fmt.Println("length of map:", len(m))
	
	// delete values from map
	delete(m, "key2")
	fmt.Println("deleted or not:", m)

	n := map[string]int{"key3":34, "key4": 45}
	fmt.Println("map n", n)
}