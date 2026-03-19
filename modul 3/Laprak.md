# <h1 align="center">Laporan Praktikum Modul 2</h1>
<p align="center">[Leonardo Farriz Garcya] - [109082530036]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}
func permutation(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
func combination(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}


```

#### soal2.go

```go
package main
import "fmt"

func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	hasila := f(g(h(a)))
	hasilb := g(h(f(b)))
	hasilc := h(f(g(c)))

	fmt.Println(hasila)
	fmt.Println(hasilb)
	fmt.Println(hasilc)
}

```

#### soal3.go

```go
package main
import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(x, y, cx, cy) <=r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
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
<p>int(bilangan bulat) </p>
<p>var: untuk mendeklarasikan variabel</p>
<p>func faktorial(n int) int :  menghitung faktorial, masukannya angka n, keluarannya juga angka, bertipe int </p>
<p>hasil := 1: membuat nilai awal "hasil" = 1</p>
<p>for i := 1; i <= n :Perulangan dari 1 sampai n</p>
<p>i++ : bertambah 1 setiap loop</p>
<p>*= : untuk perkalian lalu langsung meyimpan ke variabel yang sama</p>
<p>return hasil : mengembalikan nilai faktorial</p>
<p>func permutation(n, r int) int : fungsi permutasi dengan input n dan r, bertipe int </p>
<p>return faktorial(n) / faktorial(n-r): menghitung P(n,r) = n! / (n-r)!(Rumus)</p>
<p>func combination(n, r int) int : fungsi kombinasi dengan input n dan r, bertipe int</p>
<p>return faktorial(n) / (faktorial(r) * faktorial(n-r)): Menghitung C(n,r) = n! / (r! × (n-r)!) (Rumus)</p>
<p>a, b, c, d int : membuat 4 variabel bertipe integer</p>
<p>scan(&a, &b, &c, &d): membaca 4 input dari user, dan menyimmpannya sesuai variabel</p>
<p>Println/Print: untuk menampilkan hasil ke layar sesuai kondisi (Println membuat baris baru di akhir, Print tidak)</p>

<p>Program ini digunakan untuk menghitung permutasi dan kombinasi dari beberapa pasangan bilangan.Jadi fungsi faktorial untuk menghitung n! dengan perulangan. Kemudain ada fungsi permutation (P) dan combination (C) dengan rumus (rumus bisa di lihat di atas). Saat dijalankan, program akan membaca 4 angka inputan dari user(a,b,c,d). Kemudian menghitung P(a, c) dan C(a, c) dan menghitung P(b, d) dan C(b, d). Dan hasilnya ditampilkan dalam 2 baris output,yaitu:</p>
<p>Baris pertama, hasil dari a dan</p>
<p>Baris kedua, hasil dari b dan d</p>


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
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/output/soal%202.png" >

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
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/output/soal%203.png" >