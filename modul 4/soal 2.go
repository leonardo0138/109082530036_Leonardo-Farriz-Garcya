package main

import "fmt"

func hitungSkor(jumlahSoal, totalWaktu *int) {
	*jumlahSoal = 0
	*totalWaktu = 0

	var waktu int
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*jumlahSoal++
			*totalWaktu += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var maxSoal, minWaktu int

	pertamaKali := true

	for {
		fmt.Scan(&nama)
		if nama == "SELESAI" {
			break
		}

		var jumlahSoal, totalWaktu int
		hitungSkor(&jumlahSoal, &totalWaktu)

		if pertamaKali {
			maxSoal = jumlahSoal
			minWaktu = totalWaktu
			pemenang = nama
			pertamaKali = false
		} else if jumlahSoal > maxSoal || (jumlahSoal == maxSoal && totalWaktu < minWaktu) {

			maxSoal = jumlahSoal
			minWaktu = totalWaktu
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}
