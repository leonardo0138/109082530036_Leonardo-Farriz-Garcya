# <h1 align="center">Laporan Praktikum Modul 12 </h1>
<p align="center">[Leonardo Farriz Garcya] - [109082530036]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

func sequential(suara []int, target int) int {
	for i := 0; i < 20; i++ {
		if i+1 == target {
			return suara[i]
		}
	}
	return -1
}

func main() {
	var n int
	total := 0
	sah := 0
	data := [100]int{}
	suara := [20]int{}

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		data[total] = n
		total++
	}
	for i := 0; i < total; i++ {
		if data[i] >= 1 && data[i] <= 20 {
			sah++
			suara[data[i]-1]++
		}
	}
	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", sah)

	for i := 1; i <= 20; i++ {
		hasil := sequential(suara[:], i)
		if hasil > 0 {
			fmt.Printf("%d: %d\n", i, hasil)
		}
	}
}

```

#### soal2.go

```go
package main
import "fmt"

func ketua(suara []int) int {
	pos := 0
	for i := 1; i < 20; i++ {
		if suara[i] > suara[pos] {
			pos = i
		}
	}
	return pos
}

func wakil(suara []int, posKetua int) int {
	pos := -1
	for i := 0; i < 20; i++ {
		if i == posKetua {
			continue
		}
		if pos == -1 || suara[i] > suara[pos] {
			pos = i
		}
	}
	return pos
}
func main() {
	var n int
	total := 0
	sah := 0
	suara := [20]int{}

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		total++
		if n >= 1 && n <= 20 {
			sah++
			suara[n-1]++
		}
	}
	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", sah)
	ketua := ketua(suara[:])
	wakil := wakil(suara[:], ketua)

	fmt.Printf("Ketua RT: %d\n", ketua+1)
	fmt.Printf("Wakil ketua: %d\n", wakil+1)
}

```

#### soal3.go

```go
package main
import "fmt"

const NMAX = 1000000
var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}
func posisi(n, k int) int {
	kiri := 0
	kanan := n - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if data[tengah] == k {
			return tengah
		} else if data[tengah] < k {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)
	hasil := posisi(n, k)
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}

```
### Output Unguided :

[penjelasan]
<h2>Soal 1</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>func sequential(suara []int, target int) int: fungsi untuk mencari jumlah suara kandidat berdasarkan nomor kandidat</p>
<p>for i := 0; i < 20; i++: perulangan untuk memeriksa seluruh kandidat</p>
<p>if i+1 == target: kondisi jika nomor kandidat sama dengan target yang dicari</p>
<p>return suara[i]: mengembalikan jumlah suara kandidat tersebut</p>
<p>return -1: mengembalikan -1 jika kandidat tidak ditemukan</p>
<p>func main(): fungsi utama program</p>
<p>var n int: variabel untuk menyimpan input suara</p>
<p>total := 0: variabel untuk menghitung jumlah seluruh suara yang masuk</p>
<p>sah := 0: variabel untuk menghitung jumlah suara sah</p>
<p>data := [100]int{}: array untuk menyimpan seluruh data suara yang masuk</p>
<p>suara := [20]int{}: array untuk menyimpan jumlah suara masing-masing kandidat</p>
<p>for { }: perulangan untuk membaca data suara</p>
<p>fmt.Scan(&n): membaca nomor kandidat yang dipilih</p>
<p>if n == 0: kondisi berhenti jika input bernilai 0</p>
<p>break: menghentikan perulangan input</p>
<p>data[total] = n: menyimpan suara ke dalam array data</p>
<p>total++: menambah jumlah suara masuk</p>
<p>for i := 0; i < total; i++: perulangan untuk memeriksa seluruh suara yang masuk</p>
<p>if data[i] >= 1 && data[i] <= 20: kondisi jika suara valid (nomor kandidat 1 sampai 20)</p>
<p>sah++: menambah jumlah suara sah</p>
<p>suara[data[i]-1]++: menambah jumlah suara kandidat yang dipilih</p>
<p>fmt.Printf("Suara masuk: %d\n", total): menampilkan jumlah seluruh suara yang masuk</p>
<p>fmt.Printf("Suara sah: %d\n", sah): menampilkan jumlah suara sah</p>
<p>for i := 1; i <= 20; i++: perulangan untuk menampilkan hasil suara setiap kandidat</p>
<p>hasil := sequential(suara[:], i): memanggil fungsi untuk mendapatkan jumlah suara kandidat ke-i</p>
<p>if hasil > 0: kondisi jika kandidat memiliki suara</p>
<p>fmt.Printf("%d: %d\n", i, hasil): menampilkan nomor kandidat dan jumlah suaranya</p>

<p>Program ini digunakan untuk menghitung hasil pemungutan suara untuk 20 kandidat. Setiap suara yang masuk direpresentasikan oleh sebuah angka dari 1 sampai 20 yang menunjukkan nomor kandidat yang dipilih.</p>
<p>Program membaca data suara secara berulang hingga pengguna memasukkan angka 0 sebagai tanda akhir input. Semua data suara yang masuk disimpan ke dalam array, kemudian program menghitung jumlah seluruh suara yang diterima.</p>
<p>Selanjutnya, program memeriksa setiap suara untuk menentukan apakah suara tersebut sah. Suara dianggap sah jika nilainya berada pada rentang 1 sampai 20. Untuk setiap suara sah, program menambahkan jumlah suara kandidat yang bersangkutan dan menghitung total suara sah.</p>
<p>Program menggunakan fungsi sequential untuk mengambil jumlah suara dari kandidat tertentu berdasarkan nomor urut kandidat. Fungsi ini melakukan pencarian secara berurutan pada array hasil perhitungan suara.</p>
<p>Setelah seluruh data diproses, program menampilkan jumlah suara masuk, jumlah suara sah, serta daftar kandidat yang memperoleh suara beserta banyaknya suara yang diterima. Kandidat yang tidak memperoleh suara tidak akan ditampilkan.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%2012/output/soal%201.png" >


[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>func ketua(suara []int) int: fungsi untuk mencari kandidat dengan suara terbanyak sebagai ketua</p>
<p>pos := 0: menyimpan indeks kandidat dengan suara terbanyak sementara</p>
<p>for i := 1; i < 20; i++: perulangan untuk memeriksa seluruh kandidat</p>
<p>if suara[i] > suara[pos]: kondisi jika suara kandidat saat ini lebih banyak</p>
<p>pos = i: mengubah posisi kandidat dengan suara terbanyak</p>
<p>return pos: mengembalikan indeks kandidat yang menjadi ketua</p>
<p>func wakil(suara []int, posKetua int) int: fungsi untuk mencari kandidat dengan suara terbanyak kedua sebagai wakil ketua</p>
<p>pos := -1: inisialisasi posisi wakil ketua</p>
<p>for i := 0; i < 20; i++: perulangan untuk memeriksa seluruh kandidat</p>
<p>if i == posKetua: kondisi jika kandidat adalah ketua</p>
<p>continue: melewati kandidat ketua agar tidak dipilih lagi</p>
<p>if pos == -1 || suara[i] > suara[pos]: kondisi jika belum ada wakil atau suara kandidat lebih banyak</p>
<p>pos = i: mengubah posisi kandidat dengan suara terbanyak kedua</p>
<p>return pos: mengembalikan indeks kandidat yang menjadi wakil ketua</p>
<p>func main(): fungsi utama program</p>
<p>var n int: variabel untuk menyimpan input suara</p>
<p>total := 0: variabel untuk menghitung jumlah seluruh suara yang masuk</p>
<p>sah := 0: variabel untuk menghitung jumlah suara sah</p>
<p>suara := [20]int{}: array untuk menyimpan jumlah suara masing-masing kandidat</p>
<p>for { }: perulangan untuk membaca data suara</p>
<p>fmt.Scan(&n): membaca nomor kandidat yang dipilih</p>
<p>if n == 0: kondisi berhenti jika input bernilai 0</p>
<p>break: menghentikan perulangan input</p>
<p>total++: menambah jumlah suara masuk</p>
<p>if n >= 1 && n <= 20: kondisi jika nomor kandidat valid</p>
<p>sah++: menambah jumlah suara sah</p>
<p>suara[n-1]++: menambah jumlah suara kandidat yang dipilih</p>
<p>fmt.Printf("Suara masuk: %d\n", total): menampilkan jumlah seluruh suara yang masuk</p>
<p>fmt.Printf("Suara sah: %d\n", sah): menampilkan jumlah suara sah</p>
<p>ketua := ketua(suara[:]): memanggil fungsi untuk menentukan ketua RT</p>
<p>wakil := wakil(suara[:], ketua): memanggil fungsi untuk menentukan wakil ketua RT</p>
<p>fmt.Printf("Ketua RT: %d\n", ketua+1): menampilkan nomor kandidat yang menjadi ketua RT</p>
<p>fmt.Printf("Wakil ketua: %d\n", wakil+1): menampilkan nomor kandidat yang menjadi wakil ketua RT</p>

<p>Program ini digunakan untuk menghitung hasil pemungutan suara dan menentukan kandidat yang terpilih sebagai Ketua RT serta Wakil Ketua RT berdasarkan jumlah suara yang diperoleh.</p>
<p>Program membaca data suara secara berulang hingga pengguna memasukkan angka 0 sebagai tanda akhir input. Setiap angka yang berada pada rentang 1 sampai 20 dianggap sebagai suara sah dan dihitung sebagai suara untuk kandidat yang sesuai. Selain menghitung suara sah, program juga mencatat jumlah seluruh suara yang masuk.</p>
<p>Setelah semua suara diproses, jumlah suara setiap kandidat disimpan dalam sebuah array. Fungsi ketua digunakan untuk mencari kandidat dengan jumlah suara terbanyak, sedangkan fungsi wakil digunakan untuk mencari kandidat dengan jumlah suara terbanyak berikutnya dengan mengabaikan kandidat yang telah terpilih sebagai ketua.</p>
<p>Program kemudian menampilkan jumlah suara masuk, jumlah suara sah, nomor kandidat yang terpilih sebagai Ketua RT, dan nomor kandidat yang terpilih sebagai Wakil Ketua RT.</p>
<p>Dengan demikian, program ini menerapkan proses penghitungan suara sekaligus pencarian nilai maksimum dan maksimum kedua untuk menentukan hasil akhir pemilihan.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%2012/output/soal%202.png" >

[penjelasan]
<h2>Soal 3</h2>
<br>  
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>const NMAX = 1000000: konstanta untuk ukuran maksimum array</p>
<p>var data [NMAX]int: array global untuk menyimpan data bilangan</p>
<p>func isiArray(n int): fungsi untuk mengisi array dengan n data</p>
<p>for i := 0; i < n; i++: perulangan untuk membaca data sebanyak n elemen</p>
<p>fmt.Scan(&data[i]): membaca dan menyimpan elemen ke array</p>
<p>func posisi(n, k int) int: fungsi untuk mencari posisi nilai k menggunakan binary search</p>
<p>kiri := 0: variabel untuk menyimpan batas kiri pencarian</p>
<p>kanan := n - 1: variabel untuk menyimpan batas kanan pencarian</p>
<p>for kiri <= kanan: perulangan selama rentang pencarian masih ada</p>
<p>tengah := (kiri + kanan) / 2: menentukan indeks tengah array</p>
<p>if data[tengah] == k: kondisi jika nilai yang dicari ditemukan</p>
<p>return tengah: mengembalikan indeks tempat nilai ditemukan</p>
<p>else if data[tengah] < k: kondisi jika nilai tengah lebih kecil dari k</p>
<p>kiri = tengah + 1: menggeser batas kiri ke kanan</p>
<p>else: kondisi jika nilai tengah lebih besar dari k</p>
<p>kanan = tengah - 1: menggeser batas kanan ke kiri</p>
<p>return -1: mengembalikan -1 jika nilai tidak ditemukan</p>
<p>func main(): fungsi utama program</p>
<p>var n, k int: variabel untuk jumlah data dan nilai yang dicari</p>
<p>fmt.Scan(&n, &k): membaca jumlah data dan nilai yang akan dicari</p>
<p>isiArray(n): memanggil fungsi untuk mengisi array</p>
<p>hasil := posisi(n, k): memanggil fungsi untuk mencari posisi nilai k</p>
<p>if hasil == -1: kondisi jika nilai tidak ditemukan</p>
<p>fmt.Println("TIDAK ADA"): menampilkan bahwa nilai tidak ada dalam array</p>
<p>else: kondisi jika nilai ditemukan</p>
<p>fmt.Println(hasil): menampilkan indeks nilai yang ditemukan</p>

<p>Program ini digunakan untuk mencari posisi suatu bilangan dalam kumpulan data yang telah terurut menggunakan metode binary search. Data disimpan dalam sebuah array global dengan kapasitas hingga satu juta elemen.</p>
<p>Program terlebih dahulu membaca jumlah data (n) dan nilai yang ingin dicari (k). Selanjutnya, pengguna memasukkan n buah bilangan yang akan disimpan ke dalam array melalui fungsi isiArray.</p>
<p>Proses pencarian dilakukan oleh fungsi posisi dengan teknik binary search. Metode ini bekerja dengan membandingkan nilai yang dicari dengan elemen tengah array. Jika nilai yang dicari lebih kecil dari elemen tengah, pencarian dilanjutkan ke bagian kiri array. Sebaliknya, jika lebih besar, pencarian dilanjutkan ke bagian kanan array. Proses ini berulang hingga data ditemukan atau ruang pencarian habis.</p>
<p>Jika bilangan yang dicari ditemukan, fungsi mengembalikan indeks tempat bilangan tersebut berada. Namun, jika bilangan tidak ditemukan, fungsi mengembalikan nilai -1.</p>
<p>Pada akhir program, hasil pencarian ditampilkan. Jika data ditemukan, program menampilkan posisi atau indeksnya. Jika tidak ditemukan, program menampilkan pesan "TIDAK ADA".</p>


##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%2012/output/soal%203.png" >


