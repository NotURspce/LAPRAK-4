package main
import "fmt"

func main() {
	var bmi, tinggi, berat float64

	fmt.Print("Masukan Berat mass indeks : ")
	fmt.Scan(&bmi)

	fmt.Print("Masukan Tinggi badanmu : ")
	fmt.Scan(&tinggi)

	berat = bmi * tinggi * tinggi 
	fmt.Print("Kg : ")
	fmt.Printf("%.f", berat)

}


// program berat_badan
// kamus
//    bmi, tinggi, berat : real

// algoritma
//    input(bmi, tinggi)
//    berat <- bmi * tinggi * tinggi
//    output(berat)
// endprogram // 