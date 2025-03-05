package main

import "fmt"

func penjumlahan(x float64, z float64) float64 {
	return x + z
}

func perkalian(x float64, z float64) float64 {
	return x * z
}

func pengurangan(x float64, z float64) float64 {
	return x - z
}

func pembagian(x float64, z float64) float64 {
	return x / z
}

func main() {
	var bilangan1, bilangan2 float64
	var Opmath rune
	fmt.Println("Silakan lakukan operasi matematika (+, -, *, /) atau ketik '#' untuk keluar.")

	_, err := fmt.Scanf("%f %c %f", &bilangan1, &Opmath, &bilangan2)
	if err != nil {
		fmt.Println("Input tidak valid, pastikan formatnya benar!")
		return
	}

	for Opmath != '#' {
		switch Opmath {
		case '+':
			bilangan1 = penjumlahan(bilangan1, bilangan2)
			fmt.Println("Hasil : ", bilangan1)
		case '-':
			bilangan1 = pengurangan(bilangan1, bilangan2)
			fmt.Println("Hasil : ", bilangan1)
		case '*':
			bilangan1 = perkalian(bilangan1, bilangan2)
			fmt.Println("hasil : ", bilangan1)
		case '/':
			if bilangan2 == 0 {
				fmt.Println("input tidak valid")
			} else {
				bilangan1 = pembagian(bilangan1, bilangan2)
				fmt.Println("Hasil : ", bilangan1)
			}
		default:
			fmt.Println("Operator tidak valid")
		}
		fmt.Print("Masukkan operasi berikutnya atau '#' untuk keluar: ")

		_, err := fmt.Scanf(" %c", &Opmath)
		if err != nil {
			fmt.Println("Input tidak valid, pastikan formatnya benar!")
			return
		}

		if Opmath != '#' {
			_, err = fmt.Scanf("%f", &bilangan2)
			if err != nil {
				fmt.Println("Input tidak valid, pastikan formatnya benar!")
				return
			}
		}
	}
}
