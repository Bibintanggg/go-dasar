package main

import "fmt"

func main() {
	var point = 10

	// CONDITIONAL RENDERING BERSARNG(IF - ELSE, SWITCH CASE JADI SATU KODE)
	// disini variable di sebut duluan pada switchnya

	if point > 15 {
		switch point {
		case 10:
			fmt.Println("OKEEH SESUAII!")
		default:
			fmt.Println("NIICEE !")
		}
	} else {
		switch point {
		case 9:
			fmt.Println("BENER NIHHH GA SESUAI")
			fallthrough
		case 5:
			fmt.Println("WKWKKWK BENERAN GA SESUAI")
		default:
			fmt.Println("BISMILLAH BISA GO")
		}
	}
}
