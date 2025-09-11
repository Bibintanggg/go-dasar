package main

import "fmt"

func main() {
	var angka = [5]int{1, 2, 3, 4, 5}

	for _, i := range angka {
		fmt.Println("angka saat ini yang belum di jumlah =", i)

		penjumlahan := i + 5
		// fmt.Println("Angka yang sudah ditambah = ", penjumlahan)

		perkalian := penjumlahan * 20
		fmt.Println("Angka yang sudah dikali = ", perkalian)
	}
}
