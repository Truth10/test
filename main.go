package main

import "fmt"

func main() {
	a := [4]int{10, 40, 5, 280}
	b := [5]int{234, 5, 2, 148, 23}
	v := 42
	ok := test(a, b, v)
	if ok {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func test(a [4]int, b [5]int, v int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i]+b[j] == v {
				return true
			}
		}
	}
	return false
}
