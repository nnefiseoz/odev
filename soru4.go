package main

import "fmt"

func main() {

	//0'dan başlayıp 100'e kadar gidicek bir döngü oluşturuyoruz.
	for i := 0; i <= 100; i++ {

		//Sayıların 2 ile modunu alıyoruz.
		if i%2 == 1 {

			//Sayının modu 1 ise ekrana yazdırıyoruz.
			fmt.Println(i)
		}

	}
}
