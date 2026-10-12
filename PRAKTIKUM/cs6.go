package main
import "fmt"

func main() {
	const hargaKopi = 20000
	const pajak = 0.11

	var jumlah int
	var subtotal, total float64

	fmt.Scan(&jumlah)

	subtotal = float64(jumlah) * hargaKopi
	total = subtotal + subtotal*pajak

	fmt.Println(int(total))
}