package main
import "fmt"

func main() {
	var h, batas, tekanan float64
	const p float64 = 1025.0
	const g float64 = 9.8

	fmt.Scan(&h, &batas)
	tekanan =  p * g * h

	fmt.Println(tekanan)
	fmt.Println(tekanan <= batas)
}