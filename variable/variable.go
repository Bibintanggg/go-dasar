package main

import "fmt"

func main() {
	var namaAwal string = "Bintang"
	var namaAkhir string = "Yudha"
	var age uint8 = 17

	fmt.Println("hallo", namaAwal, namaAkhir, "umur saya", age)

	name := new(string)
	fmt.Println(*name) // menghasilkan output yang awalannya string ""
	fmt.Println(name)
}
