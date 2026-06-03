package main

import "fmt"

// Notes
// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

// var a = make(map[KeyType]ValueType)
// b := make(map[KeyType]ValueType)

func main() {
	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// a is no longer empty
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}
