package main
import "fmt"

func main() {
	var belanja, diskon, akhir int

	fmt.Print("Masukan belanja Belanjamu : ")
	fmt.Scan(&belanja)

	fmt.Print("Masukan Diskon : ")
	fmt.Scan(&diskon)

	akhir = belanja - (belanja * diskon / 100)

	fmt.Print("Total Akhir : ")
	fmt.Print(akhir)
}


// program total_belanja
// kamus
//    total, diskon, akhir : int

// algoritma
// 	  input(total, diskon)
// 	  total <- total - (total * diskon div 100)
// 	  output(akhir)
// endprogram//