package main

import "fmt"

func main() {

	fmt.Println("== FOR + RANGE TIPE STRING")
	var xs = "123"         // string
	for i, v := range xs { // for + range bertujuan untuk menggabungkan data map, slice, array dan string
		fmt.Println("Index =", i, "Value: ", v)
	}

	fmt.Println("== FOR RANGE ARRAY ==")
	var ys = [5]int{1, 2, 3, 4, 5} // for mengembalikan array sebanyak isi yang dibuat
	for i, y := range ys {
		fmt.Println("Index = ", i, "Value =", y)
	}

	fmt.Println("== FOR RANGE SLICE ==") // slice berfungsi hanya untuk menyeleksi angka yang dipilih
	var zs = ys[0:2]                     // dimulai dari 0 hingga angka ke 2
	for _, p := range zs {
		fmt.Println("Index =", p)
	}

	fmt.Println("== RANGE MAP ==")
	var kvs = map[byte]int{'a': 1, 'b': 2, 'c': 3} // single quote '' untuk karakter/ascii/byte.. byte untuk huruf map
	for i, h := range kvs {                        // int untuk angkanya
		fmt.Println("Indexnya yaitu = ", i, "Value nya adalah = ", h) // index nya dari ascii
	}

	var arraypt2 = [5]int{10, 20, 30, 40, 50}
	for index, value := range arraypt2 {
		fmt.Println("Index nya =", index, "Valuenya =", value)
	}

	var slicePt2 = arraypt2[0:3]
	for _, val := range slicePt2 {
		fmt.Println("Valuenyah = ", val)
	}

	var mapp = map[byte]int{'a': 1, 'b': 2, 'c': 3}
	for key, valueh := range mapp {
		fmt.Println("Key nya yaitu =", key, "Value =", valueh)
	}
}
