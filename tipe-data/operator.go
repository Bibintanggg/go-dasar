package main

import "fmt"

func main() {
	var aritmatika = (((2 + 3) / 2 * 3) + 10%2)
	var perbandingan = 6 >= 2

	fmt.Println(aritmatika)
	fmt.Println(perbandingan)

	var left = true
	var right = false

	var leftAndRight = left && right
	fmt.Printf("left && right \t %t", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("\nleft || right \t %t", leftOrRight)
}
