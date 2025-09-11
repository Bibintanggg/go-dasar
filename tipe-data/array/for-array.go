package main

import "fmt"

func main() {
	fruits := [5]string{"apples", "mango", "banana", "cherry", "orange"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println("Index ke = ", i, "Buahnya adalah = ", fruits[i])
	}

	println("== FOR RANGE WITH ARRAY ==")
	// fungsi _(underscore) adalah kosong/tidak dibutuhkan karena saya tulis di index.. berarti saya hanya butuh valuenya aj
	for _, y := range fruits { 
		fmt.Println( "Dan valuenya", y)
	}
}
