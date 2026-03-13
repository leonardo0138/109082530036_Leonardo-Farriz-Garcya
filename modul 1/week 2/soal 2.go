package main

import "fmt"

func main() {
	var (
		color1, color2, color3, color4 string
		temp bool
	)
	fmt.Println("Urutkan warna sesuai rules")
	temp = true
	i := 1
	for i <= 5 {
		fmt.Print("Percobaan ", i, " : ")
		fmt.Scan(&color1, &color2, &color3, &color4)
		if color1 != "merah" || color2 != "kuning" || color3 != "hijau" || color4 != "ungu" {
			temp = false
		}
		i++
	}
	fmt.Println("Berhasil :", temp)
}
