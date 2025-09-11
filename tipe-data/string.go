package main

import "fmt"

func main() {
	var message = `"hallo guysss"
	nama gua bintang
	"yudha"`

	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n") // \n untuk membuat line baru
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")

	fmt.Println(message)
}
