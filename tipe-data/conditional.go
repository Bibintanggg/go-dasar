package main

import "fmt"

func main() {
	var nilai uint8 = 80

	if nilai == 80 {
		fmt.Println("nilai kamu aman")
	} else if nilai < 80 {
		fmt.Println("harus remedial")
	} else if nilai < 50 {
		fmt.Println("mengulang mapel")
	} else if nilai > 80 {
		fmt.Println("kamu lulus")
	} else {
		fmt.Printf("maaf kamu tinggal kelas karena dibawah 80 terus %d\n", nilai)
	}

	fmt.Println("====== VAR TEMPORARY =====")
	// VAR TEMPORARY PADA IF - ELSE
	var nilaiAkhir = 87.50

	if percent := nilaiAkhir / 10; percent >= 100 {
		fmt.Printf("%% Nilai akhir kamu : %.2f\n", percent) //%.2f itu artinya 2 desimal aja yang dibutuhin
	} else if percent <= 100 {
		fmt.Printf("%% Nilai akhir kamu : %.2f\n", percent) //%.2f itu artinya 2 desimal aja yang dibutuhin
	}

	fmt.Println("\n===== SWITCH CASE ======")
	var point = 6

	switch point {
	case 5:
		fmt.Println(" NILAI SANGAT SESUAI")
	case 4:
		fmt.Println(" NILAI SANGAT SESUAI DIKIT")
	case 2:
		fmt.Println(" NILAI SANGAT TIDAK SESUAI")
		// break
	default:
		fmt.Println("GG GEMINK")
	}

	var ayam = 1

	switch ayam {
	case 5:
		fmt.Println(" AYAM YANG TIDAK SESUAI")
	case 6, 7, 8, 9:
		fmt.Println(" AYAM YANG SANGAT SESUAI")
	case 10:
		fmt.Println(" AYAM YANG SANGAT TIDAK SESUAI")
	default:
		{
			fmt.Println("GG GEMINK BROK MULTI STATEMENT")
			fmt.Println("INI MULTI STATEMENT DEFAULT")

		}
	}

	fmt.Println("\n==== SWITCH WITH IF - ELSE STYLE ====")
	var nilaiPraktek = 30

	switch {
	case nilaiPraktek == 20:
		fmt.Println("yahhh gagalll dehh okawokawoka")
	case (nilaiPraktek > 10) && (nilaiPraktek == 8):
		fmt.Println("WITH STYLE IF ELSE")
	default:
		fmt.Println("ini dengan style if else")
	}

	fmt.Println("\n=== SWITCH DENGAN FALLTHROUGH ===") // loncat 1 part ke case selanjutnya

	var rapot byte = 20

	switch {
	case rapot > 30:
		fmt.Println("Tidak sesuai")
	case rapot == 20:
		fmt.Println("Sesuai")
		fallthrough // loncat ke satu part/case selanjutnya
	case rapot < 30:
		fmt.Println("Betul sekali")
	default:
		fmt.Println("sekian untuk fallthrough")
	}

}
