package main

import "fmt"

func main() {

	//Bir sonuç değişkeni tanımlıyoruz.
	var total int

	//0'dan başlayıp 100'e kadar gidicek bir döngü oluşturuyoruz.
	for i := 0; i <= 100; i++ {

		//Sayı yüzden küçğk olduğu sürece sonuca atıyoruz
		total += i
	}

	//Toplamı yazdırıyoruz.
	fmt.Println(total)

}
