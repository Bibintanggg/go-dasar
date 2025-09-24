package main

import "fmt"

func main() {
	var names [4]string  // mendeklarasikan data sebanyak 4 data dalam array
	names[0] = "Bintang" // var names di declarate sebagai string
	names[1] = "Yudha"
	names[2] = "Putra"
	names[3] = "Purnomo"

	// [0], [1], [2], [3] adalah index sebutannya

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "orange", "mango", "banana"} // initiate array dengan horizontal

	fruits = [4]string{ // inisialisai array dengan gaya vertikal
		"apple", // koma wajib di tiap isi array dalam vertikal inisialisai
		"orange",
		"mango",
		"banana",
	}

	fmt.Println("jumlah buah sat ini ada = \t", len(fruits)) // len berfungsi untuk menghitung jumlah array
	fmt.Println("buahnya yaitu = \t", fruits)

	//array tanpa jumlah elemen
	println("== array tanpa jumlah elemen ==")
	var nums = [...]int8{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println("Total angka di nums adalah = \t", len(nums))
	fmt.Println("Angka yang tersedia yaitu = \t", nums)

	println("== array multi dimensi ==")
	number := [2][3]int8{ // 2 itu baris dan 3 itu kolom
		{2, 3, 5},
		{5, 7, 8},
	}
	fmt.Println("Array multi dimensi = \t", len(number))
	fmt.Println("Array multi dimensi = \t", number)

	println("== bikin array dengan make ==")

	cars := make([]string, 2) //make = proses membuat array
	cars[0] = "lamborghini"
	cars[1] = "aventador"

	fmt.Println("Mobil saat ini yang redi yaitu = ", cars)
}
