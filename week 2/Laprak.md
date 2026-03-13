# <h1 align="center">Laporan Praktikum Modul 2</h1>
<p align="center">[Leonardo Farriz Garcya] - [109082530036]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}


```

#### soal2.go

```go
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

```

#### soal3.go

```go
package main

import "fmt"

func main() {
	var g, kg, gr, tambahan int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&g)

	kg = g / 1000
	gr = g % 1000
	hasil := kg * 10000
	fmt.Println("Detail berat: Rp.", kg, "kg +", gr, "gr")

	if gr >= 500 && kg <= 10 {
		tambahan = gr * 5
		fmt.Println("Detail biaya: Rp.", hasil, " + Rp.", tambahan)
		hasil = hasil + tambahan
		fmt.Print("Total biaya: ", hasil)
	} else if gr < 500 && kg <= 10 {
		tambahan = gr * 15
		fmt.Println("Detail biaya: Rp.", hasil, " + Rp.", tambahan)
		hasil = hasil + tambahan
		fmt.Print("Total biaya: ", hasil)
	} else if kg > 10 {
		fmt.Println("Detail biaya: Rp.", hasil, " + Rp.", tambahan)
		fmt.Print("Total biaya: ", hasil)
	}
}


```
### Output Unguided :

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar</p>
<p>func main() : untuk titik awal eksekusi program Go, semua intruksi akan berjalan dari fungsi ini</p>
<p>var: untuk mendeklarasikan variabel</p>
<p>satu, dua, tiga, temp: variabel bertipe string(teks)</p>
<p>Println / Print: untuk menampilkan hasil atau kalimat ke layar (Println membuat baris baru di akhir, Print tidak)</p>
<p>Scanln: membaca input dari pengguna dan menyimpannya ke masing’ variabel nya</p>

<p> Jadi program ini meminta user untuk menginputkan tiga buah kata satu per satu yang disimpan ke dalam variabel satu, dua, dan tiga.
Setelah semua input dimasukkan, program akan menampilkan output awal yang berisi ketiga kata tersebut sesuai urutan input.
Kemudian program melakukan pertukaran posisi nilai variabel menggunakan variabel sementara temp. Nilai pada satu dipindahkan ke temp, lalu dua dipindahkan ke satu, tiga dipindahkan ke dua, dan nilai dari temp dipindahkan ke tiga. Proses ini membuat urutan data bergeser ke kiri.
Dan program menampilkan output akhir yang berisi urutan string yang sudah berubah.</p>

##### Output
<img src = https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/output/soal%201.png >

[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar</p>
<p>func main() : untuk titik awal eksekusi program Go, semua intruksi akan berjalan dari fungsi ini</p>
<p>var: untuk mendeklarasikan variabel</p>
<p>color1, color2, color3, color4 string: 4 variabel bertipe string(teks)</p>
<p>temp bool: variabel bertipe boolean(true/false)</p>
<p>Println/Print: untuk menampilkan hasil ke layar (Println membuat baris baru di akhir, Print tidak)</p>
<p>tempt = true: memberikan nilai awal temp yaitu true sebelum dicek </p>
<p>i := 1: untuk membuat variabel i dengan nilai awal 1</p>
<p>for i <= 5 : (for) perulangan, program akan berjalan selama i <= 5 </p>
<p>fmt.Scan(&color1, &color2, &color3, &color4): membaca 4 input dari user, dan menyimmpannya sesuai variabel</p>
<p>&: alamat variabel</p>
<p>if color1 != "merah : (if) percabangan yang digunakan untuk menjalankan kode sesuai dengan kondisi(dalam kondisi ini mengecek apakah urutan warna salah)</p>
<p>!= : operator perbandingan "Tidak sama dengan"</	p>
<p>|| : operator logika "ATAU" (OR)</p>
<p>temp = false: jika urutan warna tidak sesuai mmaka program menandai hasilnya false</p>
<p>i++ : menambah nilai i sebanyak 1</p>	

<p>Program ini meminta pengguna mengurutkan warna sesuai aturan yaitu merah, kuning, hijau, ungu. Program memberikan 5 kali percobaan untuk memasukkan empat warna tersebut. Pada setiap percobaan, pengguna memasukkan 4 warna sekaligus yang disimpan ke variabel color1, color2, color3, dan color4. Program kemudian memeriksa apakah urutan warna yang dimasukkan sesuai dengan aturan. Jika salah satu warna tidak sesuai dengan urutan yang benar, maka variabel temp akan berubah menjadi false.
Setelah semua percobaan selesai, program akan menampilkan hasil akhir "Berhasil" yang bernilai true/false </p>


##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/output/soal2.png" >

[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar</p>
<p>func main() : untuk titik awal eksekusi program Go, semua intruksi akan berjalan dari fungsi ini</p>
<p>var: untuk mendeklarasikan variabel</p>
<p>g, kg, gr, tambahan int: 3 variabel bertipe integer (bilangan bulat)</p>
<p>fmt.Scan(&g) : membaca input dari user, dan menyimmpannya ke variabel g</p>
<p>% (modulus), *(perkalian), /(pembagian)</p>
<p>if : percabangan yang digunakan untuk menjalankan kode sesuai dengan kondisi</p>
<p>else if : syarat ke 2 atau setelah if yang pertama</p>
<p>else : jika semua kondisi if tidak terpenuhi</p>
<p>&& : DAN (AND)</p>
<p>Println/Print: untuk menampilkan hasil/teks ke layar (Println membuat baris baru di akhir, Print tidak)</p>
	
<p>Program ini digunakan untuk menghitung biaya pengiriman parsel berdasarkan berat dalam gram. Pengguna diminta memasukkan berat parsel (g), kemudian program mengubahnya menjadi kilogram (kg) dan sisa gram (gr)</p>

</p>Biaya dasar dihitung dari jumlah kilogram × Rp10.000. Setelah itu program mengecek sisa gram untuk menentukan biaya tambahan:</p>
<p>Jika sisa gram ≥ 500 dan kg ≤ 10, maka biaya tambahan = gr * 5</p>
<p>Jika sisa gram < 500 dan kg ≤ 10, maka biaya tambahan = gr * 15.</p>
<p>Jika kg > 10, maka tidak ada biaya tambahan.</p>
<p>Program kemudian menampilkan detail berat, detail biaya (biaya dasar + tambahan), dan total biaya pengiriman</p>


##### Output
##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/output/soal%203.png" >