package main
import "fmt"

func main() {
	var bilangan, d1, d2, d3 int

	fmt.Print("Masukan Bilanganmu : ")
	fmt.Scan(&bilangan) // untuk menerima input angka dari user

	d1 = bilangan / 100 // input user (ratusan) dibagi 100
	d2 = bilangan % 100 / 10 // 
	d3 = bilangan %100 %10

	fmt.Print("Benar Atau Tidak : ")
	fmt.Println(d1 <= d2 && d2 <= d3)
}



