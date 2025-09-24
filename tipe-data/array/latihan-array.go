package main

import "fmt"

func main() {
	var nilai [5]int
	nilai[0] = 80
	nilai[1] = 75
	nilai[2] = 90
	nilai[3] = 90
	nilai[4] = 90

	fmt.Println("Nilai siswa seluruhnya adalah = ", nilai)

	total := 0
	for i := 0; i < len(nilai); i++ {
		total += nilai[i]
	}

	fmt.Println("Total nilai nya yaitu", total)

	// total hitung rata rata
	rataRata := float32(total) / float32(len(nilai))

	fmt.Println("Rata rata dari nilai siswa adalah = ", rataRata)
}
