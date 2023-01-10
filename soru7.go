package main

import "fmt"

var finalNotu float32
var vizeNotu float32

func main() {

	fmt.Println("Final notunuzu giriniz")
	fmt.Scanln(&finalNotu)
	fmt.Println("Vize notunuzu giriniz")
	fmt.Scanln(&vizeNotu)
	finalNotuOrtalama := (finalNotu * 70) / 100
	vizeNotuOrtalama := (vizeNotu * 30) / 100
	toplam := finalNotuOrtalama + vizeNotuOrtalama
	fmt.Println("Not ortalaması:")
	fmt.Println(toplam)
}
