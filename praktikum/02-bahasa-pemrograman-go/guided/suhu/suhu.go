package main

import "fmt"

func main() {
	var (
		celcius,kelvin,fahrenheit,raemur float64
	)


	// Input
	fmt.Println("Konversi suhu.\n Masukkan suhu dalam satuan Celcius untuk di konversi ke satuan Fahrenheit, Raemur, dan Kelvin")
	fmt.Scanln(&celcius)

	// Konversi suhu
	fahrenheit = celcius * 9 / 5 + 32
	raemur = celcius * 4 / 5
	kelvin = celcius + 273.15

	// Output
	fmt.Print("\n",fahrenheit, raemur, kelvin)
}