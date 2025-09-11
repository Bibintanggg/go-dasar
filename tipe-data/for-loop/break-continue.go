package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// loop bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, "")
		}

		fmt.Println()
	}

	for x := 4; x >= 0; x-- {
		for y := x; y < 5; y++ {
			fmt.Print(y, "")
		}

		fmt.Println("")
	}

MyLooping:
	for index := 0; index < 5; index++ {
		for value := 0; value < 5; value++ {
			if index == 1 {
				break MyLooping
			}
			fmt.Print("index [", index, "], value [", value, "]\n")
		}
	}
}
