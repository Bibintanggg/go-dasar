package main

import "fmt"

func main() {
	var a int8 = -120
	var convert float64 = float64(a)
	var stringSection string = fmt.Sprintf("%f", convert)

	fmt.Println("Konversi ke dalam int = ", a)
	fmt.Println("Konversi ke dalam float64 = ", convert)
	fmt.Println("Konversi ke dalam string = ", stringSection)
}
