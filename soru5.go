package main

import "fmt"

func main() {
	
	var arr [5]int = [5]int {9, 7, 12, 3, 5}
	enKucuk := arr[0]

	for i:=0; i<5; i++{
		if i<enKucuk{
			enKucuk=arr[i]
		}
	}
	fmt.Println(enKucuk)
}
