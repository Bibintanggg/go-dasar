package main

import "fmt"

func main() {
	const ( // deklarasi multi const
		a        = "Senin"         //inference type
		sekarang                   //inference type
		is_today bool      = false //manifest type

	)

	fmt.Println(a)
	fmt.Println(sekarang)
	fmt.Println(is_today)

	const (
		square       string = "kotak" // deklarasi mutli const
		bentuknya                     // type inference
		apakah_benar bool   = true    // manifest type
	)

	fmt.Println("=============")
	fmt.Println(square)
	fmt.Println(bentuknya)
	fmt.Println(apakah_benar)

	const satu, dua = 1, 2                     // multi const dalam 1 line // inference type
	const three, four string = "tiga", "empat" // manifest type

	const (
		matematika uint8 = 90
		bahasa           = 80
		ppkn
		fisika = 88
	)

	fmt.Println("===============")
	fmt.Println("nilai matematika = ", matematika)
	fmt.Println("nilai bahasa = ", bahasa)
	fmt.Println("nilai ppkn = ", ppkn)
	fmt.Println("nilai fisika = ", fisika)
}
