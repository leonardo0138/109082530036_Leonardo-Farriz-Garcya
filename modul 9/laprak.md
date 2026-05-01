# <h1 align="center">Laporan Praktikum Modul 9 </h1>
<p align="center">[Leonardo Farriz Garcya] - [109082530036]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}
type Lingkaran struct {
	pusat Titik
	r int
}

func jarak(p, q Titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(p, c.pusat) <= float64(c.r)
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&t.x, &t.y)

	d1 := didalam(l1, t)
	d2 := didalam(l2, t)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 && !d2 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if !d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```

#### soal2.go

```go
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

```

#### soal3.go

```go
package main

import "fmt"

func main() {
	var klubA, klubB string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var hasil []string
	pertandingan := 1
	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
		pertandingan++
	}
	for i, h := range hasil {
		fmt.Printf("Hasil %d : %s\n", i+1, h)
	}
	fmt.Println("Pertandingan selesai")
}


```

#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127
type tabel [NMAX]rune
func isiArray(t *tabel, n *int) {
	*n = 0
	var ch rune
	fmt.Scanf("%c", &ch)
	for ch != '.' && *n < NMAX {
		if ch != ' ' {
			t[*n] = ch
			*n++
		}
		fmt.Scanf("%c", &ch)
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}
func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}
func palindrom(t tabel, n int) bool {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if t[i] != t[j] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks      : ")
	isiArray(&tab, &m)

	isPalin := palindrom(tab, m)
	fmt.Print("Reverse teks : ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	fmt.Print("Palindrom ? ")
	if isPalin {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

```
### Output Unguided :

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output standar</p>
<p>import "math": digunakan untuk operasi matematika seperti akar kuadrat</p>
<p>type Titik struct {x, y int}: struktur untuk menyimpan koordinat titik (x, y)</p>
<p>type Lingkaran struct {pusat Titik; r int}: struktur untuk menyimpan data lingkaran berupa titik pusat dan jari-jari</p>
<p>func jarak(p, q Titik) float64: fungsi untuk menghitung jarak antara dua titik</p>
<p>dx := float64(p.x - q.x): menghitung selisih koordinat x dan mengubah ke float</p>
<p>dy := float64(p.y - q.y): menghitung selisih koordinat y dan mengubah ke float</p>
<p>math.Sqrt(dx*dx + dy*dy): menghitung jarak menggunakan rumus Euclidean</p>
<p>func didalam(c Lingkaran, p Titik) bool: fungsi untuk mengecek apakah titik berada di dalam lingkaran</p>
<p>jarak(p, c.pusat) <= float64(c.r): kondisi jika jarak titik ke pusat ≤ jari-jari lingkaran</p>
<p>func main(): titik awal eksekusi program</p>
<p>var l1, l2 Lingkaran: variabel untuk menyimpan dua lingkaran</p>
<p>var t Titik: variabel untuk menyimpan titik yang diuji</p>
<p>fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r): input data lingkaran 1</p>
<p>fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r): input data lingkaran 2</p>
<p>fmt.Scan(&t.x, &t.y): input koordinat titik</p>
<p>d1 := didalam(l1, t): mengecek apakah titik berada di dalam lingkaran 1</p>
<p>d2 := didalam(l2, t): mengecek apakah titik berada di dalam lingkaran 2</p>
<p>if d1 && d2: kondisi jika titik berada di dalam kedua lingkaran</p>
<p>fmt.Println("Titik di dalam lingkaran 1 dan 2"): menampilkan hasil</p>
<p>else if d1 && !d2: kondisi jika hanya di lingkaran 1</p>
<p>else if !d1 && d2: kondisi jika hanya di lingkaran 2</p>
<p>else: kondisi jika tidak di keduanya</p>
<p>fmt.Println("Titik di luar lingkaran 1 dan 2"): menampilkan hasil</p>

<p>Program ini digunakan untuk menentukan posisi sebuah titik terhadap dua lingkaran dengan memanfaatkan konsep struktur (struct) dalam Go. Terdapat dua tipe data yaitu Titik untuk merepresentasikan koordinat (x, y) dan Lingkaran yang memiliki pusat (Titik) serta jari-jari.</p>
<p>Program menerima input berupa data dua lingkaran (pusat dan jari-jari) serta satu titik yang akan diuji. Untuk menentukan posisi titik, program menggunakan fungsi jarak yang menghitung jarak antara dua titik dengan rumus Euclidean (akar dari selisih kuadrat koordinat).</p>
<p>Selanjutnya, fungsi didalam digunakan untuk mengecek apakah titik berada di dalam lingkaran, yaitu dengan membandingkan jarak titik ke pusat lingkaran dengan jari-jarinya (jarak ≤ jari-jari).</p>
<p>Hasil pengecekan disimpan dalam dua variabel boolean, lalu program menentukan output berdasarkan kondisi: apakah titik berada di dalam kedua lingkaran, hanya salah satu, atau di luar keduanya.</p>
<p>Output akhir menampilkan posisi titik sesuai hasil pengecekan tersebut.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%209/output/output%201.png" >


[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>import "math": digunakan untuk operasi matematika seperti akar kuadrat</p>
<p>func all(arr []int): fungsi untuk menampilkan semua elemen array</p>
<p>for i, v := range arr: perulangan untuk mengakses indeks dan nilai array</p>
<p>if i > 0: mengatur spasi agar rapi</p>
<p>fmt.Print(v): menampilkan elemen array</p>
<p>func ganjil(arr []int): fungsi untuk menampilkan elemen pada indeks ganjil</p>
<p>if i%2 != 0: kondisi indeks ganjil</p>
<p>func genap(arr []int): fungsi untuk menampilkan elemen pada indeks genap</p>
<p>if i%2 == 0: kondisi indeks genap</p>
<p>func kelipatanX(arr []int, x int): fungsi untuk menampilkan elemen dengan indeks kelipatan x</p>
<p>if i%x == 0: kondisi indeks merupakan kelipatan x</p>
<p>func hapusElemen(arr []int, indeks int) []int: fungsi untuk menghapus elemen pada indeks tertentu</p>
<p>append(arr[:indeks], arr[indeks+1:]...): menggabungkan array sebelum dan sesudah indeks</p>
<p>func rataRata(arr []int) float64: fungsi untuk menghitung rata-rata</p>
<p>total += v: menjumlahkan semua elemen</p>
<p>float64(total) / float64(len(arr)): menghitung rata-rata</p>
<p>func standarDeviasi(arr []int) float64: fungsi untuk menghitung standar deviasi</p>
<p>mean := rataRata(arr): menghitung rata-rata</p>
<p>diff := float64(v) - mean: selisih nilai dengan rata-rata</p>
<p>total += diff * diff: menjumlahkan kuadrat selisih</p>
<p>math.Sqrt(total / float64(len(arr))): menghitung akar dari rata-rata kuadrat selisih</p>
<p>func frekuensi(arr []int, target int) int: fungsi untuk menghitung jumlah kemunculan suatu nilai</p>
<p>if v == target: kondisi jika elemen sama dengan target</p>
<p>count++: menambah jumlah kemunculan</p>
<p>func main(): fungsi utama program</p>
<p>var n int: variabel untuk jumlah elemen array</p>
<p>fmt.Scan(&n): membaca jumlah elemen</p>
<p>arr := make([]int, n): membuat array dengan ukuran n</p>
<p>for i := 0; i < n; i++: perulangan untuk input elemen array</p>
<p>fmt.Scan(&arr[i]): membaca setiap elemen array</p>
<p>all(arr): menampilkan semua elemen</p>
<p>ganjil(arr): menampilkan elemen indeks ganjil</p>
<p>genap(arr): menampilkan elemen indeks genap</p>
<p>var x int: variabel untuk kelipatan</p>
<p>fmt.Scan(&x): membaca nilai x</p>
<p>kelipatanX(arr, x): menampilkan elemen indeks kelipatan x</p>
<p>var indeksHapus int: variabel indeks yang akan dihapus</p>
<p>fmt.Scan(&indeksHapus): membaca indeks yang akan dihapus</p>
<p>arr = hapusElemen(arr, indeksHapus): menghapus elemen pada indeks tersebut</p>
<p>all(arr): menampilkan array setelah penghapusan</p>
<p>rataRata(arr): menghitung rata-rata array</p>
<p>standarDeviasi(arr): menghitung standar deviasi array</p>
<p>var target int: variabel untuk nilai yang dicari</p>
<p>fmt.Scan(&target): membaca nilai target</p>
<p>frekuensi(arr, target): menghitung frekuensi kemunculan target</p>

<p>Program ini digunakan untuk melakukan berbagai operasi pada sebuah array bilangan bulat, mulai dari menampilkan elemen, memfilter berdasarkan indeks, hingga perhitungan statistik sederhana.</p>
<p>Program diawali dengan meminta jumlah elemen (N), kemudian pengguna menginput seluruh nilai array. Setelah itu, program menampilkan seluruh elemen, elemen dengan indeks ganjil, dan elemen dengan indeks genap.</p>
<p>Selanjutnya, program meminta sebuah nilai x untuk menampilkan elemen yang berada pada indeks kelipatan x. Program juga menyediakan fitur untuk menghapus elemen berdasarkan indeks tertentu, lalu menampilkan kembali array hasil penghapusan.</p>
<p>Setelah itu, program menghitung nilai rata-rata dan standar deviasi dari elemen array yang tersisa menggunakan perhitungan matematis. Terakhir, program menghitung berapa kali suatu angka (target) muncul dalam array (frekuensi).</p>
<p>Dengan demikian, program ini menggabungkan operasi dasar array, manipulasi data, serta analisis statistik dalam satu alur eksekusi.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%209/output/output%202.png" >

[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>func main(): fungsi utama program</p>
<p>var klubA, klubB string: variabel untuk menyimpan nama dua klub</p>
<p>fmt.Print("Klub A : "): menampilkan input untuk klub A</p>
<p>fmt.Scan(&klubA): membaca nama klub A</p>
<p>fmt.Print("Klub B : "): menampilkan input untuk klub B</p>
<p>fmt.Scan(&klubB): membaca nama klub B</p>
<p>var hasil []string: slice untuk menyimpan hasil setiap pertandingan</p>
<p>pertandingan := 1: variabel untuk nomor pertandingan</p>
<p>for { }: perulangan tak hingga untuk input skor pertandingan</p>
<p>var skorA, skorB int: variabel untuk menyimpan skor kedua klub</p>
<p>fmt.Printf("Pertandingan %d : ", pertandingan): menampilkan nomor pertandingan</p>
<p>fmt.Scan(&skorA, &skorB): membaca skor kedua klub</p>
<p>if skorA < 0 || skorB < 0: kondisi berhenti jika ada skor negatif</p>
<p>break: menghentikan perulangan</p>
<p>if skorA > skorB: kondisi jika klub A menang</p>
<p>hasil = append(hasil, klubA): menyimpan nama klub A sebagai pemenang</p>
<p>else if skorB > skorA: kondisi jika klub B menang</p>
<p>hasil = append(hasil, klubB): menyimpan nama klub B sebagai pemenang</p>
<p>else: kondisi jika skor sama</p>
<p>hasil = append(hasil, "Draw"): menyimpan hasil seri</p>
<p>pertandingan++: menambah nomor pertandingan</p>
<p>for i, h := range hasil: perulangan untuk menampilkan semua hasil</p>
<p>fmt.Printf("Hasil %d : %s\n", i+1, h): menampilkan hasil pertandingan</p>
<p>fmt.Println("Pertandingan selesai"): menampilkan bahwa pertandingan telah selesai</p>

<p>Program ini digunakan untuk mencatat hasil pertandingan antara dua klub secara berulang hingga kondisi tertentu terpenuhi. Pengguna terlebih dahulu memasukkan nama dua klub yang akan bertanding.</p>
<p>Program kemudian menjalankan perulangan untuk setiap pertandingan. Pada setiap iterasi, pengguna memasukkan skor dari kedua klub. Jika salah satu skor bernilai negatif, maka perulangan dihentikan sebagai tanda bahwa input pertandingan telah selesai.</p>
<p>Setiap hasil pertandingan akan dievaluasi: jika skor klub A lebih besar, maka klub A dianggap menang; jika skor klub B lebih besar, maka klub B menang; dan jika sama, maka hasilnya adalah draw. Hasil tersebut disimpan ke dalam sebuah array (slice).</p>
<p>Setelah semua pertandingan selesai, program menampilkan seluruh hasil pertandingan secara berurutan berdasarkan urutan input, lalu diakhiri dengan pesan bahwa pertandingan telah selesai.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%209/output/output%203.png" >


[penjelasan]
<h2>Soal 4</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>const NMAX int = 127: konstanta untuk ukuran maksimum array</p>
<p>type tabel [NMAX]rune: tipe data array berisi karakter (rune)</p>
<p>func isiArray(t *tabel, n *int): fungsi untuk mengisi array dengan karakter</p>
<p>*n = 0: inisialisasi jumlah elemen</p>
<p>var ch rune: variabel untuk menyimpan karakter input</p>
<p>fmt.Scanf("%c", &ch): membaca satu karakter</p>
<p>for ch != '.' && *n < NMAX: perulangan sampai ditemukan titik atau array penuh</p>
<p>if ch != ' ': mengabaikan spasi</p>
<p>t[*n] = ch: menyimpan karakter ke array</p>
<p>(*n)++: menambah jumlah elemen</p>
<p>fmt.Scanf("%c", &ch): membaca karakter berikutnya</p>
<p>func cetakArray(t tabel, n int): fungsi untuk menampilkan isi array</p>
<p>for i := 0; i < n; i++: perulangan untuk mencetak array</p>
<p>fmt.Printf("%c", t[i]): menampilkan karakter</p>
<p>func balikanArray(t *tabel, n int): fungsi untuk membalik isi array</p>
<p>for i, j := 0, n-1; i < j; i, j = i+1, j-1: perulangan dari depan dan belakang</p>
<p>t[i], t[j] = t[j], t[i]: menukar elemen</p>
<p>func palindrom(t tabel, n int) bool: fungsi untuk mengecek apakah array palindrom</p>
<p>for i, j := 0, n-1; i < j; i, j = i+1, j-1: membandingkan elemen depan dan belakang</p>
<p>if t[i] != t[j]: jika tidak sama maka bukan palindrom</p>
<p>return false: mengembalikan nilai salah</p>
<p>return true: jika semua sama maka palindrom</p>
<p>func main(): fungsi utama program</p>
<p>var tab tabel: variabel array karakter</p>
<p>var m int: jumlah karakter</p>
<p>fmt.Print("Teks : "): menampilkan input teks</p>
<p>isiArray(&tab, &m): mengisi array dari input</p>
<p>isPalin := palindrom(tab, m): mengecek apakah teks palindrom</p>
<p>fmt.Print("Reverse teks : "): menampilkan teks terbalik</p>
<p>balikanArray(&tab, m): membalik isi array</p>
<p>cetakArray(tab, m): mencetak array</p>
<p>fmt.Print("Palindrom ? "): menampilkan hasil pengecekan</p>
<p>if isPalin: kondisi jika palindrom</p>
<p>fmt.Println("true"): menampilkan true</p>
<p>else: kondisi jika bukan palindrom</p>
<p>fmt.Println("false"): menampilkan false</p>

<p>Program ini digunakan untuk mengolah teks yang dimasukkan oleh pengguna dan melakukan beberapa operasi seperti membalik teks serta mengecek apakah teks tersebut merupakan palindrom.</p>
<p>Input dibaca karakter per karakter menggunakan tipe data rune dan disimpan ke dalam array hingga ditemukan tanda titik (.) sebagai penanda akhir. Spasi tidak ikut disimpan, sehingga hanya karakter selain spasi yang diproses.</p>
<p>Setelah input diterima, program terlebih dahulu mengecek apakah teks tersebut adalah palindrom, yaitu apakah urutan karakter sama jika dibaca dari depan maupun dari belakang.</p>
<p>Selanjutnya, array dibalik menggunakan proses pertukaran elemen dari depan dan belakang, lalu hasil teks yang sudah dibalik ditampilkan.</p>
<p>Terakhir, program menampilkan hasil pengecekan palindrom dalam bentuk true atau false berdasarkan kondisi teks awal sebelum dibalik.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%209/output/output%204.png" >

