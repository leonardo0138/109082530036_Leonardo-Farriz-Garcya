# <h1 align="center">Laporan Praktikum Modul 3 </h1>
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
<p>fmt.Scan(&a, &b, &c, &d): membaca 4 input dari user, dan menyimmpannya sesuai variabel</p>
<p>Println/Print: untuk menampilkan hasil ke layar sesuai kondisi (Println membuat baris baru di akhir, Print tidak)</p>

<p>Program ini digunakan untuk menghitung permutasi dan kombinasi dari beberapa pasangan bilangan.Jadi fungsi faktorial untuk menghitung n! dengan perulangan. Kemudain ada fungsi permutation (P) dan combination (C) dengan rumus (rumus bisa di lihat di atas). Saat dijalankan, program akan membaca 4 angka inputan dari user(a,b,c,d). Kemudian menghitung P(a, c) dan C(a, c) dan menghitung P(b, d) dan C(b, d). Dan hasilnya ditampilkan dalam 2 baris output,yaitu:</p>
<p>Baris pertama, hasil dari a dan baris kedua, hasil dari b dan d</p>

##### Output
<img src = "https://github.com/leonardo0138/modul-3/blob/main/output/soal%201.png" >


[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar</p>
<p>func main() : untuk titik awal eksekusi program Go, semua intruksi akan berjalan dari fungsi ini</p>
<p>var: untuk mendeklarasikan variabel</p>
<p>func f/g/h(x int) int : fungsi f/g/h, dengan parameter x(menerima 1 angka x) bertipe integer dan mengembalikan integer/p>
<p> a, b, c : membuat 3 variabel bertipe integer /p>
<p>fmt.Scan(&a, &b, &c): membaca 3 input dari user, dan menyimmpannya sesuai variabel</p>
<p>&: alamat variabel</p>
<p>hasila := f(g(h(a))): memanggil fungsi secara bertingkat. Dan seterusnya begitu pula dengan hasilb,hasilc</p>
<p>Println/Print: untuk menampilkan hasil ke layar sesuai kondisi (Println membuat baris baru di akhir, Print tidak)</p>

<p>Program ini digunakan untuk menghitung hasil dari komposisi beberapa fungsi matematika sederhana dengan urutan tertentu Usere diminta memasukkan tiga bilangan yaitu a, b, dan c. Program memiliki tiga fungsi:</p>
<p>f(x) = x² (kuadrat)<p>
<p>g(x) = x − 2<p>
<p>h(x) = x + 1<p>
<p>Setelah menerima input, program tidak langsung menghitung secara terpisah, tetapi menggunakan komposisi fungsi, yaitu hasil dari satu fungsi akan menjadi input untuk fungsi berikutnya<p>
<p>Setelah itu, program menghitung tiga hasil:</p>
<p>f(g(h(a))) → nilai a diproses berurutan oleh fungsi h, lalu g, lalu f </p> 
<p>g(h(f(b))) → nilai b diproses oleh f, lalu h, lalu g </p>
<p>h(f(g(c))) → nilai c diproses oleh g, lalu f, lalu h </p>
<p>Setiap hasil disimpan dalam variabel berbeda, lalu program menampilkannya dalam tiga baris output sesuai urutan perhitungan<p>

##### Output
<img src = "https://github.com/leonardo0138/modul-3/blob/main/output/soal%202.png" >

[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar</p>
<p>math : untuk fungsi matematika (Sqrt)</p>
<p>func main() : untuk titik awal eksekusi program Go, semua intruksi akan berjalan dari fungsi ini</p>
<p>var: untuk mendeklarasikan variabel</p>
<p>func jarak(a, b, c, d float64) float64: fungsi untuk menghitung jarak di dua titik, menerima 4 nilai float64 dan mengembalikan float64 </p>
<p>func didalam(cx, cy, r, x, y float64) bool: fungsi untuk mengecek apakah titik berada di dalam lingkaran, menerima nilai float64(bilangan real) dan mengembalikan boolean (true/false)</p>
<p>variabel pada soal no 3 ini dideklarasikan dengan tipe float64(bilangan real)</p>
<p>fmt.Scan: membaca input dari user, dan menyimmpannya sesuai variabel. </p>
<p>&: alamat variabel</p>
<p>dalam1 := didalam(cx1, cy1, r1, x, y) : memanggil fungsi didalam untuk mengecek apakah titik berada di dalam lingkaran 1, sekaligus mendeklarasikan variabel baru dalam1 dan menyimpan hasilnya, begitu pula dengan "dalam2"(perbedaannya ngecek lingkaran 2)</p>
<p>	if : percabangan yang digunakan untuk menjalankan kode sesuai dengan kondisi</p>
<p>else if : syarat ke 2 atau setelah if yang pertama</p>
<p>else : jika semua kondisi if tidak terpenuhi</p>
<p>&& : DAN (AND)</p>
<p>Println/Print: untuk	 menampilkan hasil/teks ke layar (Println membuat baris baru di akhir, Print tidak)</p>

<p>Program ini digunakan untuk mengecek apakah sebuah titik berada di dalam atau di luar dua lingkaran.Pengguna diminta memasukkan data:
<p>-Titik pusat dan jari-jari lingkaran 1</p>
<p>-Titik pusat dan jari-jari lingkaran 2</p>
<p>-Satu titik yang akan dicek posisinya</p>
<p>Program kemudian menghitung jarak titik ke masing-masing pusat lingkaran menggunakan rumus jarak. Setelah itu, jarak tersebut dibandingkan dengan jari-jari lingkaran untuk menentukan apakah titik berada di dalam atau di luar</p>
<p>Hasil pengecekan dibagi menjadi beberapa kondisi:</p>
<p>-Jika titik berada di kedua lingkaran, maka output: Titik di dalam lingkaran 1 dan 2</p>
<p>-Jika hanya di lingkaran 1, maka output: titik di dalam lingkaran 1</p>
<p>-Jika hanya di lingkaran 2, maka output: titik di dalam lingkaran 2</p>
<p>-Jika tidak berada di keduanya, maka output: titik di luar lingkaran 1 dan 2 </p>


##### Output
<img src = "https://github.com/leonardo0138/modul-3/blob/main/output/soal%203.png" >