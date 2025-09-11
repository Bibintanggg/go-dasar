package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89     //jangakauan int
	var negativeNumber int16 = -31273 //jangkauan int

	var positive byte = 50       //positive.. byte = uint8
	var negative rune = -3237123 //negative..  rune = int32

	fmt.Println("Positive number = ", positiveNumber)
	fmt.Println("Negative number = ", negativeNumber)

	fmt.Println("Positive = ", positive)
	fmt.Println("Positive = ", negative)

	fmt.Printf("positive number %d", positive) 

}
