package main

import "fmt"

func main() {

	//En küçük elemanı bulmak için bir array listesi hazırlıyoruz.
	var arr [10]int = [10]int{20, 4, 5, 7, 11, 13, 2, 24, 65, 21}

	//En küçük elemanı atamak için bir değişken oluşturuyoruz.
	minNum := arr[0]

	//Listeyi baştan sona kontrol etmesi için bir döngü oluşturuyoruz.
	for i := 0; i < len(arr); i++ {

		//En küçük elemanı arıyoruz.
		if arr[i] < minNum {
			minNum = arr[i]

		}
	}
	//En küçük elemanı yazdırıyoruz.
	fmt.Println(minNum)

}
