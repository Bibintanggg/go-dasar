package main

import "fmt"

func main() {
	var nilaiUjian float64 = 75.20
	var kkm float64 = 80.00

	if passCheck := nilaiUjian; passCheck >= kkm {
		fmt.Printf("%% Lulus %.2f\n", passCheck)
	} else {
		fmt.Println("Tidak lulus")
	}

}
