# <h1 align="center">Laporan Praktikum Modul 4 </h1>
<p align="center">[Leonardo Farriz Garcya] - [109082530036]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d, pr ,co int
	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &pr)
	combination(a, c, &co)
	fmt.Println(pr, co)
	
	permutation(b, d, &pr)
	combination(b, d, &co)
	fmt.Println(pr, co)
}

```

#### soal2.go

```go
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


```

#### soal3.go

```go
package main

import "fmt"

func deret(n int) {
	for {
		fmt.Print(n, " ")
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	deret(n)
}

```
### Output Unguided :

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar</p>
<p>func main() : untuk titik awal eksekusi program Go, semua intruksi akan berjalan dari fungsi ini</p>
<p>int(bilangan bulat)</p>
<p>var: untuk mendeklarasikan variabel</p>
<p>func factorial(n int, hasil *int) : Menghitung nilai faktorial dari n (n!) dan hasil *int (pointer untuk menyimpan hasil) </p>
<p>for i := 1; i <= n :Perulangan dari 1 sampai n</p>
<p>i++ : bertambah 1 setiap loop</p>
<p>*hasil = 1: pointer hasil, nilai di alamat tersebut berubah</p>
<p>func permutation(n, r int, hasil *int) : fungsi permutation menghitung nilai menggunakan rumus P(n,r)=n!/(n−r)!, dengan n dan r sebagai parameter. Nilai faktorial disimpan dalam variabel fn dan fnr, kemudian hasil perhitungan disimpan melalui pointer hasil. </p>
<p>func combination(n, r int, hasil *int): fungsi combination menghitung nilai kombinasi menggunakan rumus 
C(n,r)=n!/r!(n−r)!, dengan n dan r sebagai parameter. Nilai faktorial disimpan dalam variabel fn, fr, dan fnr, kemudian hasil perhitungan disimpan melalui pointer hasil. </p>

<p>fmt.Scan(&a, &b, &c, &d): membaca 4 input dari user, dan menyimmpannya sesuai variabel.</p>
<p> permutation(a, c, &pr): menghitung P(a,c) dan hasil disimpan ke variabel pr</p>
<p>	combination(a, c, &co):  menghitung C(c,a) Hasil disimpan ke variabel co</p>
<p> fmt.Println(pr, co):menampilkan niali pr dan co</p>
<p>begitu juga dengan permutation dan combination berikutnya</p>
<p>Println/Print: untuk menampilkan hasil ke layar sesuai kondisi (Println membuat baris baru di akhir, Print tidak)</p>

<p>Program ini digunakan untuk menghitung permutasi dan kombinasi dari beberapa pasangan bilangan. Program memiliki fungsi factorial untuk menghitung nilai n! dengan perulangan dan hasilnya disimpan menggunakan pointer. Terdapat fungsi permutation (P) dan combination (C) yang menggunakan rumus matematika permutasi dan kombinasi.</p>
<p>Saat dijalankan, program akan membaca 4 angka input dari user (a, b, c, d). Kemudian program menghitung P(a, c) dan C(a, c), lalu menghitung P(b, d) dan C(b, d). Setiap hasil perhitungan disimpan melalui parameter pointer sehingga langsung mengubah nilai variabel tujuan.</p>
<p>Hasil akhirnya ditampilkan dalam 2 baris output, yaitu baris pertama berisi hasil dari pasangan a dan c, dan baris kedua berisi hasil dari pasangan b dan d.</p>

##### Output
<img src = "https://github.com/leonardo0138/modul-3/blob/main/output/soal%201.png" >


[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar</p>
<p>func main() : untuk titik awal eksekusi program Go, semua intruksi akan berjalan dari fungsi ini</p>
<p>var: untuk mendeklarasikan variabel</p>
<p>func hitungSkor(jumlahSoal, totalWaktu *int) : mendefinisikan fungsi dengan 2 paramenter pointer yaitu jumah soal dan totalwaktu bertipe integer dan menginisialisasi pointer jumlahsoal dan total waktu = nilai 0 (	*jumlahSoal/	*totalWaktu = 0) </p>
<p>for i := 0; i < 8; i++ : perulangan(loop) sebanyak 8 kali</p>
<p>fmt.Scan(&waktu): membaca inputan dari user dan disimpan ke variabelnya</p>
<p>&: alamat variabel</p>
<p>if waktu < 301 : cek apakah waktu ≤ 301 menit (valid)</p>
<p>(*jumlahSoal)++: menambah jumlah soal yang berhasil</p>
<p>*totalWaktu += waktu: menambahkan waktu ke total waktu</p>
<p>var nama, pemenang string: untuk mendeklarasikan variabel tersebut bertipe string</p>
<p>var maxSoal, minWaktu int: untuk mendeklarasikan variabel tersebut bertipe int</p>
<p>pertamaKali := true: penanda untuk peserta pertama</p>
<p>for: perulangan disini digunakan untuk membaca data peserta</p>
<p>if nama == "SELESAI": kondisi untuk menghentikan input</p>
<p>break: keluar dari loop jika selesai</p>
<p>var jumlahSoal, totalWaktu int: selain mendeklarasikan variabel, varialbel untuk menyimpan hasil perhitungan peserta</p>
<p>hitungSkor(&jumlahSoal, &totalWaktu): memanggil fungsi untuk menghitung skor peserta</p>
<p>if pertamaKali: inisialisasi data pemenang pertama</p>
<p>maxSoal = jumlahSoal: untuk menyimpan jumlah soal tertinggi sementara</p>
<p>minWaktu = totalWaktu: untuk menyimpan waktu tercepat sementara</p>
<p>pemenang = nama: untuk menyimpan nama pemenang sementara</p>
<p>pertamaKali = false: menandakan bahwa data pertama sudah diproses</p>
<p>else if jumlahSoal > maxSoal || (jumlahSoal == maxSoal && totalWaktu < minWaktu): kondisi untuk menentukan pemenang baru</p>
<p>|| : atau, && : AND</P>
<p> = :memberi nilai, == membandingkan nilai(T/F)</p>
<p>maxSoal = jumlahSoal: update jumlah soal terbaik</p>
<p>minWaktu = totalWaktu: update waktu terbaik</p>
<p>pemenang = nama: update nama pemenang</p>
<p>Println/Print: untuk menampilkan hasil ke layar sesuai kondisi (Println membuat baris baru di akhir, Print tidak)</p>

<p>Program ini digunakan untuk menentukan pemenang berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan. Program memiliki fungsi hitungSkor yang bertugas menghitung berapa banyak soal yang diselesaikan serta total waktu yang dibutuhkan, dengan menggunakan parameter pointer agar hasilnya langsung tersimpan ke variabel utama.</p>
<p>Di dalam fungsi hitungSkor, program akan membaca 8 input waktu pengerjaan. Jika waktu kurang dari 301 menit, maka soal dianggap berhasil, sehingga jumlahSoal bertambah dan totalWaktu diakumulasikan. Jika tidak, soal tersebut tidak dihitung.</p>
<p>Saat dijalankan, program akan terus membaca nama peserta sampai input "SELESAI". Untuk setiap peserta, program menghitung jumlah soal yang berhasil dan total waktunya, lalu membandingkan dengan peserta sebelumnya.</p>
<p>Pemenang ditentukan berdasarkan jumlah soal terbanyak. Jika jumlah soal sama, maka dipilih peserta dengan total waktu paling kecil. Hasil akhirnya ditampilkan berupa nama pemenang, jumlah soal yang diselesaikan, dan total waktu pengerjaan.</p>



##### Output
<img src = "https://github.com/leonardo0138/modul-3/blob/main/output/soal%202.png" >

[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar</p>
<p>func main() : untuk titik awal eksekusi program Go, semua intruksi akan berjalan dari fungsi ini</p>
<p>var: untuk mendeklarasikan variabel</p>
<p>func deret(n int): fungsi untuk menghasilkan deret berdasarkan aturan tertentu</p>
<p>for { }: perulangan tak hingga sampai kondisi berhenti terpenuhi</p>
<p>fmt.Print(n, " "): menampilkan nilai n</p>
<p>if n == 1: kondisi berhenti jika nilai sudah 1</p>
<p>break: menghentikan perulangan</p>
<p>if n%2 == 0: kondisi jika n adalah bilangan genap</p>
<p>n = n / 2: jika genap maka n dibagi 2</p>
<p>else: jika n adalah bilangan ganjil</p>
<p>n = 3*n + 1: jika ganjil maka n dikali 3 lalu ditambah 1</p>
<p>fmt.Scan(&n): membaca input nilai n dari user</p>
<p>deret(n): memanggil fungsi untuk menampilkan deret angka</p>
<p>Program ini digunakan untuk menampilkan deret bilangan berdasarkan aturan tertentu (dikenal sebagai pola Collatz). Program menerima satu input bilangan bulat n dari pengguna, kemudian memprosesnya melalui fungsi deret.</p>
<p>Di dalam fungsi deret, program akan mencetak nilai n secara berulang hingga mencapai angka 1. Setiap langkah mengikuti aturan, jika n adalah bilangan genap, maka dibagi 2.
Jika n adalah bilangan ganjil, maka dikalikan 3 lalu ditambah 1.</p>
<p>Proses ini berlangsung terus dalam perulangan hingga n bernilai 1, dan setiap nilai ditampilkan secara berurutan dalam satu baris output.</p>
<p>Dengan demikian, program ini menunjukkan bagaimana suatu bilangan dapat berubah mengikuti aturan tertentu sampai mencapai kondisi berhenti.</p>

##### Output
<img src = "https://github.com/leonardo0138/modul-3/blob/main/output/soal%203.png" >