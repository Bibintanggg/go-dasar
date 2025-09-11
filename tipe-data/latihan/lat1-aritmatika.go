package main

import "fmt"

func main() {
	// a := 10
	// b := 3 //var uint = 3

	var a uint8 = 10
	var b uint8 = 3

	penjumlahan := a + b
	pengurangan := a - b
	perkalian := a * b
	pembagian := a / b

	fmt.Println("hasil dari penjumlahan = ", penjumlahan)
	fmt.Println("hasil dari pengurangan = ", pengurangan)
	fmt.Println("hasil dari perkalian = ", perkalian)
	fmt.Println("hasil dari pembagian = ", pembagian)

	fmt.Println("cara dari penjumlahan 10 + 3 =", a+b)
	fmt.Println("cara dari pengurangan 10 - 3 =", a-b)
	fmt.Println("cara dari perkalian 10 * 3 =", a*b)
	fmt.Println("cara dari pembagian 10 / 3 =", a/b)

}
