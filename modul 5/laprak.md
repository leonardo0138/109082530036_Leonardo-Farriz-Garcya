# <h1 align="center">Laporan Praktikum Modul 4 </h1>
<p align="center">[Leonardo Farriz Garcya] - [109082530036]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}

```

#### soal2.go

```go
package main

import "fmt"

func bintang(n int) {
	if n == 0 {
		return
	}
	bintang(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n)
}


```

#### soal3.go

```go
package main

import "fmt"

func faktor(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}


```

#### soal4.go

```go
package main
import "fmt"	
func baris(n int) {
	if n == 0 {
		return
	}
	 fmt.Printf("%d ", n)
	baris(n - 1)
	if n != 1 {
		fmt.Printf("%d ", n)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	baris(n)
}




```

#### soal5.go

```go
package main

import "fmt"
func ganjil (n int) {
	if n == 0 {
		return
	}
	ganjil (n - 1)
	if n % 2 != 0 {
		fmt.Printf("%d ", n)
	}
}
func main() {	
	var n int	
	fmt.Scan(&n)
	ganjil (n)
}	

```

#### soal6.go

```go
package main

import "fmt"

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(pangkat(x, y))
}


```
### Output Unguided :

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar (cetak ke layar)</p>
<p>func fibonacci(n int) int: mendefinisikan fungsi rekursif dengan parameter n bertipe integer dan mengembalikan integer</p>
<p>if n == 0: kondisi dasar, jika n = 0 maka hasilnya 0</p>
<p>return 0: mengembalikan nilai 0</p>
<p>else if n == 1: kondisi dasar kedua, jika n = 1</p>
<p>return 1: mengembalikan nilai 1</p>
<p>return fibonacci(n-1) + fibonacci(n-2): rumus rekursif Fibonacci, menjumlahkan dua nilai sebelumnya</p>
<p>func main(): titik awal eksekusi program</p>
<p>for i := 0; i <= 10; i++: perulangan dari 0 sampai 10</p>
<p>fmt.Printf("%d ", fibonacci(i)): mencetak hasil Fibonacci dari i dengan format integer dan spasi</p>
<p>fibonacci(i): pemanggilan fungsi rekursif</p>
<p>i++: penambahan nilai i setiap iterasi</p>
<p>%d: format untuk menampilkan angka integer</p>

<p>Program ini digunakan untuk menampilkan deret Fibonacci dari indeks 0 hingga 10 dengan memanfaatkan fungsi rekursif fibonacci. Setiap nilai dihitung berdasarkan penjumlahan dua nilai sebelumnya.</p>
<p>Pada fungsi fibonacci, terdapat kondisi dasar yaitu n = 0 menghasilkan 0 dan n = 1 menghasilkan 1. Untuk n > 1, fungsi akan memanggil dirinya sendiri dengan parameter (n-1) dan (n-2), lalu menjumlahkan hasilnya.</p>
<p>Fungsi main melakukan perulangan dari 0 sampai 10, memanggil fibonacci(i) pada setiap iterasi, kemudian mencetak hasilnya secara berurutan.
Proses rekursi berlangsung hingga mencapai kondisi dasar, lalu hasil dikembalikan bertahap hingga diperoleh nilai akhir. Output yang dihasilkan adalah deret Fibonacci: 0 1 1 2 3 5 8 13 21 34 55.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%205/output/output%201.png" >


[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar (cetak ke layar)</p>
<p>func bintang(n int): mendefinisikan fungsi dengan parameter n bertipe integer</p>
<p>func bintang(n int): mendefinisikan fungsi dengan parameter n bertipe integer</p>
<p>if n == 0: kondisi dasar rekursi, jika n = 0 maka berhenti</p>
<p>return: keluar dari fungsi</p>
<p>bintang(n - 1): pemanggilan fungsi secara rekursif dengan nilai n dikurangi 1</p>
<p>for i := 0; i < n; i++: perulangan untuk mencetak bintang sebanyak n kali</p>
<p>fmt.Print("*"): mencetak karakter "*" tanpa pindah baris</p>
<p>fmt.Println(): pindah ke baris baru setelah mencetak bintang</p>
<p>func main(): fungsi utama (awal program)</p>
<p>var n int: deklarasi variabel n bertipe integer</p>
<p>fmt.Scan(&n): membaca input dari user dan disimpan ke variabel n</p>
<p>bintang(n): memanggil fungsi bintang untuk mencetak pola</p>
<p>rekursi: fungsi memanggil dirinya sendiri sampai kondisi berhenti</p>
<p>&n: alamat memori variabel n </p>
<p>i++: penambahan nilai i setiap perulangan</p>

<p>Program ini digunakan untuk menampilkan pola bintang berbentuk segitiga menggunakan pendekatan rekursif.</p>
<p>Pada fungsi bintang, terdapat kondisi dasar yaitu ketika n = 0, fungsi berhenti. Jika tidak, fungsi akan memanggil dirinya sendiri dengan parameter (n-1), sehingga proses berjalan dari nilai terkecil ke terbesar.</p>
<p>Setelah pemanggilan rekursif, program mencetak n buah karakter "*" menggunakan perulangan, lalu pindah ke baris berikutnya. Hal ini menyebabkan pola tercetak secara bertahap dari sedikit ke banyak.</p>
<p>Fungsi main membaca input n, kemudian memanggil fungsi bintang(n) untuk menghasilkan pola. Output yang dihasilkan adalah segitiga bintang dengan jumlah baris sesuai nilai n.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%205/output/output%202.png" >

[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar (cetak ke layar)</p>
<p>func faktor(n int, i int): mendefinisikan fungsi rekursif dengan parameter n (bilangan) dan i (pembagi awal)</p>
<p>if i > n: kondisi berhenti, jika i sudah lebih besar dari n</p>
<p>return: keluar dari fungsi</p>
<p>if n % i == 0: mengecek apakah i adalah faktor dari n (habis dibagi)</p>
<p>fmt.Printf("%d ", i): mencetak nilai i (faktor dari n)</p>
<p>faktor(n, i+1): pemanggilan rekursif dengan menaikkan nilai i</p>
<p>func main(): fungsi utama program</p>
<p>var n int: deklarasi variabel n bertipe integer</p>
<p>fmt.Scan(&n): membaca input dari user</p>
<p>faktor(n, 1): memanggil fungsi faktor dimulai dari i = 1</p>
<p>rekursi: fungsi memanggil dirinya sendiri sampai kondisi i > n</p>
<p>n % i == 0: operasi modulus(%) untuk mencari faktor</p>
<p>i+1: menaikkan nilai pembagi secara bertahap</p>
<p>%d:  format untuk menampilkan angka integer</p>

<p>Program ini digunakan untuk menampilkan faktor-faktor dari suatu bilangan dengan memanfaatkan fungsi rekursif bernama faktor. Fungsi ini memeriksa setiap bilangan dari 1 hingga n untuk menentukan apakah merupakan faktor.</p>
<p>Pada fungsi faktor, kondisi dasar terjadi ketika nilai i lebih besar dari n, sehingga proses rekursi dihentikan. Jika tidak, program akan mengecek apakah n habis dibagi i (n % i == 0). Jika ya, maka nilai i dicetak sebagai faktor.</p>
<p>Selanjutnya, fungsi akan memanggil dirinya sendiri dengan nilai i+1, sehingga proses pengecekan dilakukan secara berurutan dari 1 hingga n.Fungsi main membaca input n, kemudian memanggil faktor(n, 1) untuk memulai proses pencarian faktor. Output yang dihasilkan adalah seluruh faktor dari bilangan n.</p>
##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%205/output/output%203.png" >


[penjelasan]
<h2>Soal 4</h2>
<br>
<p>func baris(n int): mendefinisikan fungsi rekursif dengan parameter n</p>
<p>if n == 0: kondisi berhenti, jika n = 0 maka fungsi selesai</p>
<p>return: keluar dari fungsi</p>
<p>fmt.Printf("%d ", n): mencetak nilai n (bagian menurun)</p>
<p>baris(n - 1): pemanggilan rekursif dengan n dikurangi 1</p>
<p>if n != 1: kondisi agar angka 1 tidak dicetak dua kali</p>
<p>fmt.Printf("%d ", n): mencetak kembali nilai n (bagian menaik)</p>
<p>func main(): fungsi utama program</p>
<p>var n int: deklarasi variabel n bertipe integer</p>
<p>fmt.Scan(&n): membaca input dari user</p>
<p>baris(n): memanggil fungsi baris</p>
<p>&: alamat variabel</p>
<p>%d:  format untuk menampilkan angka integer</p>

<p>Program ini digunakan untuk menampilkan pola bilangan menurun lalu menaik (simetris) menggunakan rekursi melalui fungsi baris. Pola dimulai dari n hingga 1, lalu kembali naik hingga n.</p>
<p>Pada fungsi baris, kondisi dasar terjadi ketika n = 0, sehingga rekursi dihentikan. Sebelum pemanggilan rekursif, nilai n dicetak sehingga menghasilkan urutan menurun.</p>
<p>Fungsi kemudian memanggil dirinya sendiri dengan parameter (n-1). Setelah rekursi selesai, program akan mencetak kembali nilai n selama n tidak sama dengan 1, sehingga terbentuk urutan menaik tanpa mengulang angka tengah.</p>
<p>Fungsi main membaca input n dan memanggil baris(n). Output yang dihasilkan adalah pola bilangan simetris, misalnya untuk n = 5 menghasilkan 5 4 3 2 1 2 3 4 5.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%205/output/output%204.png" >


[penjelasan]
<h2>Soal 5</h2>
<br>
<p>func ganjil(n int): mendefinisikan fungsi rekursif dengan parameter n</p>
<p>if n == 0: kondisi berhenti, jika n = 0 maka fungsi selesai</p>
<p>return: keluar dari fungsi</p>
<p>ganjil(n - 1): pemanggilan rekursif dengan n dikurangi 1 (menuju nilai terkecil)</p>
<p>if n % 2 != 0: mengecek apakah n adalah bilangan ganjil</p>
<p>fmt.Printf("%d ", n): mencetak nilai n jika ganjil</p>
<p>func main(): fungsi utama program</p>
<p>var n int: deklarasi variabel n bertipe integer</p>
<p>fmt.Scan(&n): membaca input dari user</p>
<p>ganjil(n): memanggil fungsi untuk mencetak bilangan ganjil</p>
<p>%: modulus</p>
<p>%d:  format untuk menampilkan angka integer</p>
<p>&: alamat variabel</p>

<p>Program ini digunakan untuk menampilkan bilangan ganjil dari 1 hingga n dengan memanfaatkan fungsi rekursif bernama ganjil. Program akan memeriksa setiap bilangan secara berurutan dan hanya mencetak yang memenuhi kondisi ganjil.</p>
<p>Pada fungsi ganjil, kondisi dasar terjadi ketika n = 0, sehingga rekursi dihentikan. Fungsi terlebih dahulu memanggil dirinya sendiri dengan parameter (n-1) agar proses berjalan dari nilai kecil ke besar.</p>
<p>Setelah pemanggilan rekursif, program mengecek apakah n adalah bilangan ganjil (n % 2 != 0). Jika ya, maka nilai n dicetak.Fungsi main membaca input n dan memanggil ganjil(n). Output yang dihasilkan adalah deret bilangan ganjil dari 1 sampai n.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%205/output/output%205.png" >


[penjelasan]
<h2>Soal 6</h2>
<br>
<p>func pangkat(x int, y int) int: mendefinisikan fungsi rekursif untuk menghitung x pangkat y</p>
<p>if y == 0: kondisi dasar, jika pangkat = 0</p>
<p>return 1: hasilnya 1 (karena x⁰ = 1)</p>
<p>return x * pangkat(x, y-1): rumus rekursif, mengalikan x dengan hasil pangkat sebelumnya</p>
<p>func main(): fungsi utama program</p>
<p>var x, y int: deklarasi variabel x dan y bertipe integer</p>
<p>fmt.Scan(&x, &y): membaca input x dan y dari user</p>
<p>fmt.Println(pangkat(x, y)): mencetak hasil pemanggilan fungsi pangkat</p>
<p>&: alamat variabel</p>

<p>Program ini digunakan untuk menghitung hasil perpangkatan suatu bilangan. Fungsi ini menghitung nilai x^y dengan mengalikan x secara berulang sebanyak y kali.</p>
<p>Pada fungsi pangkat, kondisi dasar terjadi ketika y = 0, sehingga fungsi mengembalikan nilai 1. Jika y lebih dari 0, maka fungsi akan mengembalikan hasil perkalian x dengan pemanggilan rekursif pangkat(x, y-1).</p>
<p>Proses rekursi berlangsung hingga nilai y mencapai 0, kemudian hasil dikembalikan secara bertahap hingga diperoleh nilai akhir perpangkatan.Fungsi main membaca input x dan y, lalu memanggil pangkat(x, y). Output yang dihasilkan adalah hasil dari x pangkat y.</p>
##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%205/output/output%206.png" >