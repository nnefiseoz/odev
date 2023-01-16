package main

import "fmt"

func main() {

	var arr []int = []int{32, 46, 70, 46, 80, 73, 80, 70, 45, 32}

	var n map[int]int = make(map[int]int)
	for _, i := range arr {
		n[i] = n[i] + 1
	}
	fmt.Println(n)
}
