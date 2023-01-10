package main

import "fmt"

func main() {

	arr := []int{5, 5, 5, 10, 5}
	toplam := 0
	for i := 0; i < len(arr); i++ {
		toplam += arr[i]
	}
	ortalama := (float64(toplam)) / 2
	fmt.Println("Toplam = ", toplam, "Ortalma = ", ortalama)
}
