package main

import "fmt"

func main() {
	var height float32
	var weight float32
	fmt.Println("Please enter your height and weight")
	fmt.Scanf("%f %f", &height, &weight)

	weight /= (height * height)

	if weight < 18.5 {
		fmt.Println("Thin")
	} else if weight > 18.5 && weight < 24.9 {
		fmt.Println("Normal")
	} else if weight > 25 && weight < 29.9 {
		fmt.Println("Overweight")
	} else if weight > 30 && weight < 34.9 {
		fmt.Println("1st degree obese")
	} else if weight > 35 && weight < 39.9 {
		fmt.Println("2st degree obese")
	} else if weight > 40 {
		fmt.Println("3st degree obese")
	} else {
		fmt.Println("Please enter a valid Height and Weight")
	}
}
