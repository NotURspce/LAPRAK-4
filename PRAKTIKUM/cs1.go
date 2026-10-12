package main
import "fmt"

func main() {
	var detik, jam, menit int
	fmt.Scan(&detik)
	jam = detik / 3600
	menit = (detik % 3600) / 60
	detik = detik % 60
	fmt.Println(jam, "jam", menit, "menit dan", detik, "detik")
}


// PSEUDOCODE

//ALGORITMA KonversiWaktu

//Kamus
//Detik, Menit, Jam : int

//PROGRAM
//input(Detik)
//Jam <- Detik div 3600
//Menit <- (Detik mod 3600) div 60
//Detik <- Detik mod 60
//output(jam, menit, detik)
//ENDPROGRAM

