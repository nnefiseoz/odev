package main

import "fmt"

func main() {

	var limit int
	fmt.Println("Dizinin eleman limitini giriniz")
	fmt.Scan(&limit)

	slice := make([]int, limit)

	for i := 0; i < limit; i++ {
		var n int
		fmt.Println("dizinin", i, "'ninci elemanını yazınız")
		fmt.Scan(&n)
		slice[i] = n

	}
	for i := e; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				var n int
				n = slice[i]
				slice[i] = slice[j]
				slice[j] = n
			}
		}
	}
	fmt.Println("dizinizin büyükten küçüğe sıralanmış hali :", slice)

}
