package main

import (
	"fmt"
)

func main() {
	// elem := []string{"hello", "table", "chair", "code"}
	// elemMap := make(map[int]string)
	// for i := 0; i < len(elem); i++ {
	// 	elemMap[i] = elem[i]
	// }
	// fmt.Println(elemMap)

	const (
		_ = iota
		c = iota
	)

	// elem := []string{"hello", "table", "chair", "code"}
	// elemMap := make(map[string]int)
	// for i := 0; i < len(elem); i++ {
	// 	elemMap[elem[i]] = i
	// }
	// fmt.Println(elemMap)

	// if val, ok := dict["foo"]; ok {
	// 	//do something here
	// }

	elem := []string{"hello", "table", "chair", "hello", "hello", "code"}
	elemMap := make(map[string]int)
	for _, v := range elem {
		// returns returnValue, which contains value in elemMap[v] and status, which is true if status has a value and false otherwise
		if returnValue, status := elemMap[v]; status {
			elemMap[v] = returnValue + 1
		} else {
			elemMap[v] = c
		}

	}
	fmt.Println(elemMap)
	fmt.Printf("%T \t %v \n", elemMap, elemMap)
}
