package main

import "fmt"

func main() {

	var evenOrOdd int

	//Kullanıcıdan bir sayı alıyoruz.
	fmt.Println("Enter a number!")

	fmt.Scanln(&evenOrOdd)
	
	//Aldığımız sayının 2 ile modunu alıyoruz 
	if evenOrOdd%2 == 0 {

		//Işlemin sonucu 0'a eşitse ekrana sayının çift olduğunu yazdırıyoruz.
		fmt.Println("Even number")

		//Işlemin sonucu 0'a eşit değilse sayının tek olduğunu ekrana yazdırıyoruz.	
	} else {
		fmt.Println("Odd number")
	}

}
