# <h1 align="center">Laporan Praktikum Modul 10 </h1>
<p align="center">[Leonardo Farriz Garcya] - [109082530036]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

const N = 1000
func hitung(berat[N] float64,n int, min *float64, max *float64){
	*min = berat[0]
	*max= berat[0]

	for  i:= 1;  i<n; i++ {
		if berat[i]<*min {
			*min= berat[i]
		}
		if berat[i]>*max{
			*max= berat[i]
		}
	}
}
func main (){
	var berat[N] float64
	var max, min float64
	var n int
	fmt.Scan(&n)

	for i:= 0; i<n; i++{
		fmt.Scan(&berat[i])
	}
	hitung(berat, n, &min,&max)
	fmt.Printf("%.2f %.2f\n",min, max )
}

```

#### soal2.go

```go
package main

import "fmt"

const N = 1000

func main() {
	var berat [N]float64
	var x, y int
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	w := (x + y - 1) / y
	var tot [N]float64

	for i := 0; i < x; i++ {
		idx := i / y
		tot[idx] += berat[i]
	}

	for i := 0; i < w; i++ {
		fmt.Printf("%.2f", tot[i])
		if i < w-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	for i := 0; i < w; i++ {
		nIkan := y
		if i == w-1 && x%y != 0 {
			nIkan = x % y
		}
		rata := tot[i] / float64(nIkan)
		fmt.Printf("%.2f", rata)
		if i < w-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

```

#### soal3.go

```go
package main 
import "fmt"

type balita [100] float64

func hitung(bb balita, n int, min *float64, max *float64){
	*min= bb[0]
	*max= bb[0]

	for  i:= 1;  i<n; i++ {
		if bb[i]<*min {
			*min= bb[i]
		}

		if bb[i]>*max{
			*max= bb[i]
		}
	}
}
func rerata (bb balita, n int)float64{
	total := 0.0
	for i:=0; i<n;i++{
		total += bb[i]
	}
	return total/float64(n)
}
func main(){
	var bb balita
	var n int
	var min, max float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i:= 0; i<n;i++{
		fmt.Print("masukan berat balita ke - ",i+1," : ")
		fmt.Scan(&bb[i])
	}
	hitung(bb,n,&min,&max)
	fmt.Printf("Berat balita minimum: %.2f kg \n",min)
	fmt.Printf("Berat balita maximum: %.2f kg \n",max)
	fmt.Printf("Rerata berat balita: %.2f kg", rerata(bb,n))
}

```
### Output Unguided :

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>  
<p>import "fmt": digunakan untuk input/output</p>
<p>const N = 1000: konstanta untuk ukuran maksimum array</p>
<p>func hitung(berat [N]float64, n int, min *float64, max *float64): fungsi untuk mencari nilai minimum dan maksimum pada array</p>
<p>*min = berat[0]: inisialisasi nilai minimum dengan elemen pertama array</p>
<p>*max = berat[0]: inisialisasi nilai maksimum dengan elemen pertama array</p>
<p>for i := 1; i < n; i++: perulangan untuk memeriksa semua elemen array mulai dari indeks 1</p>
<p>if berat[i] < *min: kondisi jika elemen lebih kecil dari minimum</p>
<p>*min = berat[i]: mengubah nilai minimum</p>
<p>if berat[i] > *max: kondisi jika elemen lebih besar dari maksimum</p>
<p>*max = berat[i]: mengubah nilai maksimum</p>
<p>func main(): fungsi utama program</p>
<p>var berat [N]float64: array untuk menyimpan data berat</p>
<p>var max, min float64: variabel untuk menyimpan nilai maksimum dan minimum</p>
<p>var n int: variabel jumlah data</p>
<p>fmt.Scan(&n): membaca jumlah data</p>
<p>for i := 0; i < n; i++: perulangan untuk input data berat</p>
<p>fmt.Scan(&berat[i]): membaca setiap elemen array</p>
<p>hitung(berat, n, &min, &max): memanggil fungsi untuk mencari nilai minimum dan maksimum</p>
<p>fmt.Printf("%.2f %.2f\n", min, max): menampilkan nilai minimum dan maksimum dengan 2 angka di belakang koma</p>

<p>Program ini digunakan untuk mencari nilai berat minimum dan maksimum dari sejumlah data berat badan yang dimasukkan pengguna. Data berat disimpan dalam sebuah array bertipe float64 dengan kapasitas maksimal 1000 elemen.</p>
<p>Program terlebih dahulu membaca jumlah data (n), kemudian menerima input nilai berat sebanyak n buah dan menyimpannya ke dalam array. Setelah itu, program memanggil fungsi hitung untuk menentukan nilai terkecil dan terbesar dari data tersebut.</p>
<p>Di dalam fungsi hitung, nilai awal minimum dan maksimum diisi menggunakan elemen pertama array. Selanjutnya, program melakukan perulangan mulai dari indeks kedua hingga akhir data untuk membandingkan setiap elemen.</p>
<p>Jika ditemukan nilai yang lebih kecil dari minimum, maka nilai minimum diperbarui. Sebaliknya, jika ditemukan nilai yang lebih besar dari maksimum, maka nilai maksimum juga diperbarui.</p>
<p>Hasil akhir berupa berat minimum dan maksimum kemudian ditampilkan dengan format dua angka di belakang koma.</p>



##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%209/output/output%201.png" >


[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>const N = 1000: konstanta untuk ukuran maksimum array</p>
<p>func main(): fungsi utama program</p>
<p>var berat [N]float64: array untuk menyimpan data berat ikan</p>
<p>var x, y int: variabel untuk jumlah data dan jumlah ikan per kelompok</p>
<p>fmt.Scan(&x, &y): membaca jumlah data dan jumlah ikan per kelompok</p>
<p>for i := 0; i < x; i++: perulangan untuk input data berat ikan</p>
<p>fmt.Scan(&berat[i]): membaca setiap berat ikan</p>
<p>w := (x + y - 1) / y: menghitung jumlah kelompok</p>
<p>var tot [N]float64: array untuk menyimpan total berat tiap kelompok</p>
<p>for i := 0; i < x; i++: perulangan untuk menghitung total berat tiap kelompok</p>
<p>idx := i / y: menentukan indeks kelompok</p>
<p>tot[idx] += berat[i]: menambahkan berat ikan ke kelompok</p>
<p>for i := 0; i < w; i++: perulangan untuk menampilkan total berat tiap kelompok</p>
<p>fmt.Printf("%.2f", tot[i]): menampilkan total berat dengan 2 angka di belakang koma</p>
<p>if i < w-1: kondisi agar tidak ada spasi di akhir output</p>
<p>fmt.Print(" "): menampilkan spasi antar output</p>
<p>for i := 0; i < w; i++: perulangan untuk menghitung rata-rata tiap kelompok</p>
<p>nIkan := y: jumlah ikan tiap kelompok diinisialisasi sebanyak y</p>
<p>if i == w-1 && x%y != 0: kondisi jika kelompok terakhir jumlah ikannya tidak penuh</p>
<p>nIkan = x % y: mengubah jumlah ikan pada kelompok terakhir</p>
<p>rata := tot[i] / float64(nIkan): menghitung rata-rata berat ikan tiap kelompok</p>
<p>fmt.Printf("%.2f", rata): menampilkan rata-rata dengan 2 angka di belakang koma</p>
<p>fmt.Println(): pindah ke baris baru</p>

<p>Program ini digunakan untuk mengelompokkan data berat ikan dan menghitung total serta rata-rata berat pada setiap kelompok. Data berat disimpan dalam array bertipe float64 dengan kapasitas maksimal 1000 elemen.</p>
<p>Program terlebih dahulu membaca dua input, yaitu jumlah seluruh data berat ikan (x) dan jumlah ikan dalam setiap kelompok (y). Setelah itu, pengguna memasukkan nilai berat ikan satu per satu ke dalam array.</p>
<p>Selanjutnya, program menghitung berapa banyak kelompok yang diperlukan menggunakan pembagian pembulatan ke atas. Kemudian dilakukan proses pengelompokan, di mana setiap berat ikan dijumlahkan ke kelompok tertentu berdasarkan indeksnya.</p>
<p>Setelah semua total kelompok diperoleh, program menampilkan jumlah berat masing-masing kelompok dengan format dua angka di belakang koma.</p>
<p>Program kemudian menghitung rata-rata berat setiap kelompok dengan membagi total berat kelompok dengan jumlah ikan pada kelompok tersebut. Jika kelompok terakhir tidak penuh, maka rata-ratanya dihitung berdasarkan jumlah ikan yang sebenarnya.</p>
<p>Hasil akhir yang ditampilkan adalah total berat tiap kelompok dan rata-rata berat tiap kelompok secara berurutan.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%209/output/output%202.png" >

[penjelasan]
<h2>Soal 3</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>type balita [100] float64: tipe data array untuk menyimpan berat balita</p>
<p>func hitung(bb balita, n int, min *float64, max *float64): fungsi untuk mencari berat minimum dan maksimum</p>
<p>*min = bb[0]: inisialisasi berat minimum dengan elemen pertama array</p>
<p>*max = bb[0]: inisialisasi berat maksimum dengan elemen pertama array</p>
<p>for i := 1; i < n; i++: perulangan untuk memeriksa seluruh data berat balita</p>
<p>if bb[i] < *min: kondisi jika berat lebih kecil dari minimum</p>
<p>*min = bb[i]: mengubah nilai minimum</p>
<p>if bb[i] > *max: kondisi jika berat lebih besar dari maksimum</p>
<p>*max = bb[i]: mengubah nilai maksimum</p>
<p>func rerata(bb balita, n int) float64: fungsi untuk menghitung rata-rata berat balita</p>
<p>total := 0.0: variabel untuk menyimpan total berat</p>
<p>for i := 0; i < n; i++: perulangan untuk menjumlahkan semua berat</p>
<p>total += bb[i]: menambahkan berat balita ke total</p>
<p>return total / float64(n): mengembalikan hasil rata-rata</p>
<p>func main(): fungsi utama program</p>
<p>var bb balita: variabel array untuk menyimpan data berat balita</p>
<p>var n int: variabel jumlah data balita</p>
<p>var min, max float64: variabel untuk menyimpan berat minimum dan maksimum</p>
<p>fmt.Print("Masukan banyak data berat balita : "): menampilkan input jumlah data</p>
<p>fmt.Scan(&n): membaca jumlah data balita</p>
<p>for i := 0; i < n; i++: perulangan untuk input data berat balita</p>
<p>fmt.Print("masukan berat balita ke - ", i+1, " : "): menampilkan urutan input balita</p>
<p>fmt.Scan(&bb[i]): membaca berat balita</p>
<p>hitung(bb, n, &min, &max): memanggil fungsi untuk mencari berat minimum dan maksimum</p>
<p>fmt.Printf("Berat balita minimum: %.2f kg \n", min): menampilkan berat minimum</p>
<p>fmt.Printf("Berat balita maximum: %.2f kg \n", max): menampilkan berat maksimum</p>
<p>fmt.Printf("Rerata berat balita: %.2f kg", rerata(bb, n)): menampilkan rata-rata berat balita</p>

<p>Program ini digunakan untuk mengolah data berat badan balita, yaitu mencari berat minimum, berat maksimum, dan menghitung rata-rata berat badan dari sejumlah data yang dimasukkan pengguna.</p>
<p>Program menggunakan tipe data array bernama balita untuk menyimpan maksimal 100 data berat badan bertipe float64. Pengguna terlebih dahulu memasukkan jumlah data balita, kemudian memasukkan berat badan setiap balita satu per satu.</p>
<p>Fungsi hitung digunakan untuk menentukan berat minimum dan maksimum. Nilai awal minimum dan maksimum diambil dari elemen pertama array, lalu program melakukan perulangan untuk membandingkan seluruh data berat lainnya.</p>
<p>Jika ditemukan berat yang lebih kecil, maka nilai minimum diperbarui. Sebaliknya, jika ditemukan berat yang lebih besar, maka nilai maksimum diperbarui.</p>
<p>Selain itu, program juga memiliki fungsi rerata yang digunakan untuk menghitung rata-rata berat balita. Fungsi ini menjumlahkan seluruh data berat, kemudian membaginya dengan jumlah data yang dimasukkan.</p>
<p>Hasil akhir yang ditampilkan adalah berat minimum, berat maksimum, dan rata-rata berat balita dengan format dua angka di belakang koma.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%209/output/output%203.png" >


