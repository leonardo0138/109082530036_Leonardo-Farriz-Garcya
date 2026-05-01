package main

import (
	"fmt"
	"math"
)

func all(arr []int) {
	fmt.Print("Semua elemen: ")
	for i, v := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
func ganjil(arr []int) {
	fmt.Print("Indeks ganjil: ")
	for i, v := range arr {
		if i%2 != 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}

func genap(arr []int) {
	fmt.Print("Indeks genap: ")
	for i, v := range arr {
		if i%2 == 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}

func kelipatanX(arr []int, x int) {
	fmt.Printf("Indeks kelipatan %d: ", x)
	for i, v := range arr {
		if i%x == 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}

func hapusElemen(arr []int, indeks int) []int {
	return append(arr[:indeks], arr[indeks+1:]...)
}
func rataRata(arr []int) float64 {
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

func standarDeviasi(arr []int) float64 {
	mean := rataRata(arr)
	total := 0.0
	for _, v := range arr {
		diff := float64(v) - mean
		total += diff * diff
	}
	return math.Sqrt(total / float64(len(arr)))
}

func frekuensi(arr []int, target int) int {
	count := 0
	for _, v := range arr {
		if v == target {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}
	all(arr)
	ganjil(arr)
	genap(arr)

	var x int
	fmt.Print("Masukkan nilai x untuk kelipatan: ")
	fmt.Scan(&x)
	kelipatanX(arr, x)

	var indeksHapus int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeksHapus)
	arr = hapusElemen(arr, indeksHapus)
	fmt.Print("Array setelah hapus: ")

	all(arr)
	fmt.Printf("Rata-rata: %.2f\n", rataRata(arr))
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi(arr))

	var target int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&target)
	fmt.Printf("Frekuensi %d: %d\n", target, frekuensi(arr, target))
}