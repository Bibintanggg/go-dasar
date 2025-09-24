package main

import "fmt"

func main() {
	var fruits [4]int
	fruits[0] = 10 // apel
	fruits[1] = 5  // pisang
	fruits[2] = 8  // mangga
	fruits[3] = 12 // jeruk

	total := 0
	for index, value := range fruits {
		fmt.Println("Kotak ", index, "Berisi ", value)
	}
	for i := 0; i < len(fruits); i++ {
		total += fruits[i]
	}

	fmt.Println("\nTotal dari buah ", total)

	// dikurang
	if fruits - 2 {
		
	}
}
