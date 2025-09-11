package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka nya ada", i)
	}

	//for dengan pengkondisian
	println("\n=== FOR DENGAN PENGKONDISIAN ===")
	var a = 0

	for a < 5 {
		fmt.Println("Angka saat ini: ", a)
		a++
	}

	//for tanpa pengkondisian
	fmt.Println("\n=== FOR TANPA KONDISI ===")
	var b = 0

	for {
		fmt.Println("ANGKA SEKARANG =", b)

		b++
		if b == 5 {
			break
		}
	}
}
