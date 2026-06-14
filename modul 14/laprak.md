# <h1 align="center">Laporan Praktikum Modul 14 </h1>
<p align="center">[Leonardo Farriz Garcya] - [109082530036]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

const NMAX = 1000000
var arr [NMAX]int

func selectionSort(n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		temp := arr[i]
		arr[i] = arr[idxMin]
		arr[idxMin] = temp
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}
		selectionSort(m)

		for j := 0; j < m; j++ {
			fmt.Print(arr[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

```

#### soal2.go

```go
package main
import "fmt"

const NMAX = 1000000

var arrGanjil [NMAX]int
var arrGenap [NMAX]int

func sortGanjil(n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arrGanjil[j] < arrGanjil[idxMin] {
				idxMin = j
			}
		}
		temp := arrGanjil[i]
		arrGanjil[i] = arrGanjil[idxMin]
		arrGanjil[idxMin] = temp
	}
}

func sortGenap(n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if arrGenap[j] > arrGenap[idxMax] {
				idxMax = j
			}
		}
		temp := arrGenap[i]
		arrGenap[i] = arrGenap[idxMax]
		arrGenap[idxMax] = temp
	}
}

func main() {
	var n, m, num int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		var nGanjil int = 0
		var nGenap int = 0

		for j := 0; j < m; j++ {
			fmt.Scan(&num)
			if num%2 != 0 {
				arrGanjil[nGanjil] = num
				nGanjil++
			} else {
				arrGenap[nGenap] = num
				nGenap++
			}
		}

		sortGanjil(nGanjil)
		sortGenap(nGenap)

		var pertama int = 1

		for j := 0; j < nGanjil; j++ {
			if pertama == 1 {
				fmt.Print(arrGanjil[j])
				pertama = 0
			} else {
				fmt.Print(" ", arrGanjil[j])
			}
		}

		for j := 0; j < nGenap; j++ {
			if pertama == 1 {
				fmt.Print(arrGenap[j])
				pertama = 0
			} else {
				fmt.Print(" ", arrGenap[j])
			}
		}
		fmt.Println()
	}
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

#### soal4.go

```go
package main
import "fmt"

const NMAX = 1000

type tabInt [NMAX]int

func insertionSort(arr *tabInt, n int) {
	for i := 1; i < n; i++ {
		key := (*arr)[i]
		j := i - 1
		for j >= 0 && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}
}

func jarak(arr tabInt, n int) {
	if n > 1 {
		diff := arr[1] - arr[0]
		status := 1

		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != diff {
				status = 0
			}
		}

		if status == 1 {
			fmt.Printf("Data berjarak %d\n", diff)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}

func main() {
	var arr tabInt
	var n int
	var num int

	fmt.Scan(&num)
	for num >= 0 {
		arr[n] = num
		n++
		fmt.Scan(&num)
	}

	insertionSort(&arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	jarak(arr, n)
}

```


#### soal5.go

```go
package main
import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax + 1]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i <= n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func cetakFavorit(pustaka DaftarBuku, n int) {
	if n > 0 {
		maxRating := pustaka[1].rating
		for i := 2; i <= n; i++ {
			if pustaka[i].rating > maxRating {
				maxRating = pustaka[i].rating
			}
		}

		for i := 1; i <= n; i++ {
			if pustaka[i].rating == maxRating {
				fmt.Printf("%s, %s, %s, %d\n", pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun)
			}
		}
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 2; i <= n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 1 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 1; i <= limit; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 1
	kanan := n
	found := -1

	for kiri <= kanan && found == -1 {
		med := (kiri + kanan) / 2
		if pustaka[med].rating == r {
			found = med
		} else if pustaka[med].rating < r {
			kanan = med - 1
		} else {
			kiri = med + 1
		}
	}

	if found != -1 {
		fmt.Printf("%s, %s, %s, %d, %d, %d\n", pustaka[found].judul, pustaka[found].penulis, pustaka[found].penerbit, pustaka[found].tahun, pustaka[found].eksemplar, pustaka[found].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}
func main() {
	var n, targetRating int
	var pustaka DaftarBuku

	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	fmt.Scan(&targetRating)

	cetakFavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, targetRating)
}


```
### Output Unguided :

[penjelasan]
<h2>Soal 1</h2>
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

<p>Ini digunakan untuk mencari posisi suatu bilangan dalam sebuah array yang telah terurut menggunakan algoritma binary search. Array disimpan secara global dengan kapasitas maksimal satu juta elemen sehingga dapat menampung data dalam jumlah besar.</p>
<p>Program diawali dengan membaca jumlah elemen (n) dan nilai yang akan dicari (k). Setelah itu, pengguna memasukkan n buah bilangan yang disimpan ke dalam array melalui fungsi isiArray.</p>
<p>Pencarian dilakukan oleh fungsi posisi dengan memanfaatkan dua batas pencarian, yaitu bagian kiri dan kanan array. Pada setiap langkah, program menghitung indeks tengah dan membandingkan nilainya dengan bilangan yang dicari. Jika nilai tengah sama dengan k, maka indeks tersebut langsung dikembalikan. Jika nilai tengah lebih kecil dari k, pencarian dilanjutkan ke bagian kanan. Sebaliknya, jika nilai tengah lebih besar dari k, pencarian dilanjutkan ke bagian kiri.</p>
<p>Proses ini terus berlangsung hingga data ditemukan atau tidak ada lagi bagian array yang dapat diperiksa. Jika data ditemukan, fungsi mengembalikan indeks tempat data berada. Jika tidak ditemukan, fungsi mengembalikan nilai -1.</p>
<p>Jika data ditemukan, program menampilkan posisi indeksnya. Jika data tidak ditemukan, program menampilkan pesan "TIDAK ADA".</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%2014/output/soal%201.png" >


[penjelasan]
<h2>Soal 2</h2>
<br>
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>const NMAX = 1000000: konstanta untuk ukuran maksimum array</p>
<p>var arrGanjil [NMAX]int: array global untuk menyimpan bilangan ganjil</p>
<p>var arrGenap [NMAX]int: array global untuk menyimpan bilangan genap</p>
<p>func sortGanjil(n int): fungsi untuk mengurutkan bilangan ganjil secara menaik (ascending)</p>
<p>for i := 0; i < n-1; i++: perulangan untuk proses selection sort</p>
<p>idxMin := i: menyimpan indeks nilai terkecil sementara</p>
<p>for j := i + 1; j < n; j++: mencari nilai terkecil pada bagian array yang belum terurut</p>
<p>if arrGanjil[j] < arrGanjil[idxMin]: kondisi jika ditemukan nilai yang lebih kecil</p>
<p>idxMin = j: memperbarui indeks nilai terkecil</p>
<p>temp := arrGanjil[i]: menyimpan nilai sementara untuk pertukaran</p>
<p>arrGanjil[i] = arrGanjil[idxMin]: memindahkan nilai terkecil ke posisi yang sesuai</p>
<p>arrGanjil[idxMin] = temp: menyelesaikan proses pertukaran nilai</p>
<p>func sortGenap(n int): fungsi untuk mengurutkan bilangan genap secara menurun (descending)</p>
<p>for i := 0; i < n-1; i++: perulangan untuk proses selection sort</p>
<p>idxMax := i: menyimpan indeks nilai terbesar sementara</p>
<p>for j := i + 1; j < n; j++: mencari nilai terbesar pada bagian array yang belum terurut</p>
<p>if arrGenap[j] > arrGenap[idxMax]: kondisi jika ditemukan nilai yang lebih besar</p>
<p>idxMax = j: memperbarui indeks nilai terbesar</p>
<p>temp := arrGenap[i]: menyimpan nilai sementara untuk pertukaran</p>
<p>arrGenap[i] = arrGenap[idxMax]: memindahkan nilai terbesar ke posisi yang sesuai</p>
<p>arrGenap[idxMax] = temp: menyelesaikan proses pertukaran nilai</p>
<p>func main(): fungsi utama program</p>
<p>var n, m, num int: variabel untuk jumlah kasus, jumlah data, dan nilai bilangan</p>
<p>fmt.Scan(&n): membaca jumlah kasus yang akan diproses</p>
<p>for i := 0; i < n; i++: perulangan untuk memproses setiap kasus</p>
<p>fmt.Scan(&m): membaca jumlah bilangan pada kasus tersebut</p>
<p>var nGanjil int = 0: penghitung jumlah bilangan ganjil</p>
<p>var nGenap int = 0: penghitung jumlah bilangan genap</p>
<p>for j := 0; j < m; j++: perulangan untuk membaca setiap bilangan</p>
<p>fmt.Scan(&num): membaca sebuah bilangan</p>
<p>if num%2 != 0: kondisi jika bilangan ganjil</p>
<p>arrGanjil[nGanjil] = num: menyimpan bilangan ganjil ke array ganjil</p>
<p>nGanjil++: menambah jumlah bilangan ganjil</p>
<p>else: kondisi jika bilangan genap</p>
<p>arrGenap[nGenap] = num: menyimpan bilangan genap ke array genap</p>
<p>nGenap++: menambah jumlah bilangan genap</p>
<p>sortGanjil(nGanjil): mengurutkan bilangan ganjil secara menaik</p>
<p>sortGenap(nGenap): mengurutkan bilangan genap secara menurun</p>
<p>var pertama int = 1: penanda untuk mengatur format output agar tidak diawali spasi</p>
<p>for j := 0; j < nGanjil; j++: perulangan untuk menampilkan bilangan ganjil yang sudah diurutkan</p>
<p>if pertama == 1: kondisi untuk mencetak elemen pertama tanpa spasi</p>
<p>fmt.Print(arrGanjil[j]): menampilkan bilangan ganjil</p>
<p>pertama = 0: menandakan bahwa elemen pertama sudah dicetak</p>
<p>else: kondisi untuk elemen berikutnya</p>
<p>fmt.Print(" ", arrGanjil[j]): menampilkan spasi dan bilangan ganjil</p>
<p>for j := 0; j < nGenap; j++: perulangan untuk menampilkan bilangan genap yang sudah diurutkan</p>
<p>if pertama == 1: kondisi jika belum ada elemen yang dicetak</p
<p>fmt.Print(arrGenap[j]): menampilkan bilangan genap</p>
<p>pertama = 0: menandakan bahwa elemen pertama sudah dicetak</p>
<p>else: kondisi untuk elemen berikutnya</p>
<p>fmt.Print(" ", arrGenap[j]): menampilkan spasi dan bilangan genap</p>
<p>fmt.Println(): pindah ke baris baru setelah semua bilangan ditampilkan</p>

<p>Program ini digunakan untuk mengelompokkan sejumlah bilangan menjadi bilangan ganjil dan bilangan genap, kemudian mengurutkannya dengan aturan yang berbeda sebelum ditampilkan kembali.</p>
<p>Program terlebih dahulu membaca jumlah kasus uji (n). Pada setiap kasus, program membaca banyaknya bilangan (m), kemudian menerima m buah bilangan yang akan diproses. Setiap bilangan diperiksa dan dipisahkan ke dalam dua array berbeda, yaitu array bilangan ganjil dan array bilangan genap.</p>
<p>Setelah proses pemisahan selesai, bilangan ganjil diurutkan menggunakan algoritma selection sort secara menaik (dari nilai terkecil ke terbesar). Sementara itu, bilangan genap juga diurutkan menggunakan selection sort, tetapi secara menurun (dari nilai terbesar ke terkecil).</p>
<p>Setelah kedua kelompok selesai diurutkan, program menampilkan seluruh bilangan ganjil terlebih dahulu sesuai urutan menaik, kemudian dilanjutkan dengan seluruh bilangan genap sesuai urutan menurun. Hasil setiap kasus ditampilkan dalam satu baris.</p>
<p>Ini menggabungkan proses klasifikasi data berdasarkan paritas bilangan (ganjil dan genap), pengurutan dengan dua kriteria berbeda, serta penggabungan hasil untuk menghasilkan susunan akhir yang sesuai dengan ketentuan.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%2014/output/soal%202.png" >

[penjelasan]
<h2>Soal 3</h2>
<br>  
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>const NMAX = 1000000: konstanta untuk ukuran maksimum array</p>
<p>var arr [NMAX]int: array global untuk menyimpan data bilangan</p>
<p>func insertionSort(n int): fungsi untuk mengurutkan data menggunakan metode insertion sort secara menaik</p>
<p>for i := 1; i < n; i++: perulangan untuk memproses setiap elemen mulai dari indeks 1</p>
<p>key := arr[i]: menyimpan nilai yang akan disisipkan ke posisi yang tepat</p>
<p>j := i - 1: menyimpan indeks elemen sebelum key</p>
<p>for j >= 0 && arr[j] > key: perulangan selama elemen sebelumnya lebih besar dari key</p>
<p>arr[j+1] = arr[j]: menggeser elemen ke kanan</p>
<p>j--: berpindah ke indeks sebelumnya</p>
<p>arr[j+1] = key: menempatkan key pada posisi yang sesuai</p>
<p>func main(): fungsi utama program</p>
<p>var n int = 0: variabel untuk menghitung jumlah data yang tersimpan</p>
<p>var num int: variabel untuk menyimpan bilangan yang dibaca</p>
<p>fmt.Scan(&num): membaca bilangan dari input</p>
<p>for num != -5313: perulangan selama input bukan -5313</p>
<p>if num == 0: kondisi jika input bernilai 0</p>
<p>if n > 0: kondisi jika sudah ada data yang tersimpan</p>
<p>insertionSort(n): mengurutkan seluruh data yang tersimpan</p>
<p>if n%2 != 0: kondisi jika jumlah data ganjil</p>
<p>fmt.Println(arr[n/2]): menampilkan nilai tengah (median)</p>
<p>else: kondisi jika jumlah data genap</p>
<p>fmt.Println((arr[n/2-1] + arr[n/2]) / 2): menampilkan rata-rata dua nilai tengah sebagai median</p>
<p>else: kondisi jika input bukan 0</p>
<p>arr[n] = num: menyimpan bilangan ke dalam array</p>
<p>n++: menambah jumlah data yang tersimpan</p>
<p>fmt.Scan(&num): membaca bilangan berikutnya</p>


<p>Program ini digunakan untuk menyimpan sekumpulan bilangan, mengurutkannya, dan menampilkan nilai median setiap kali menerima perintah tertentu. Data disimpan dalam sebuah array dengan kapasitas maksimal satu juta elemen.</p>
<p>Program membaca bilangan secara berulang hingga ditemukan nilai khusus -5313 yang menandakan akhir input. Setiap bilangan selain 0 dan -5313 akan disimpan ke dalam array sebagai bagian dari kumpulan data.</p>
<p>Ketika program menerima nilai 0, seluruh data yang telah tersimpan akan diurutkan menggunakan algoritma insertion sort secara menaik. Setelah data terurut, program menghitung median dari kumpulan data tersebut.</p>
<p>Jika jumlah data ganjil, median adalah elemen yang berada tepat di tengah array. Jika jumlah data genap, median diperoleh dari rata-rata dua elemen yang berada di tengah array menggunakan pembagian bilangan bulat.</p>
<p>Setiap kali nilai 0 dimasukkan, program akan menampilkan median berdasarkan seluruh data yang telah tersimpan sampai saat itu. Proses ini berlanjut hingga pengguna memasukkan nilai -5313 untuk mengakhiri program.</p>
<p>Program menggabungkan proses penyimpanan data, pengurutan menggunakan insertion sort, dan perhitungan median secara bertahap berdasarkan input yang diberikan.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%2014/output/soal%203.png" >

[penjelasan]
<h2>Soal 4</h2>
<br>  
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>const NMAX = 1000: konstanta untuk ukuran maksimum array</p>
<p>type tabInt [NMAX]int: tipe data array integer dengan kapasitas NMAX elemen</p>
<p>func insertionSort(arr *tabInt, n int): fungsi untuk mengurutkan array secara menaik menggunakan Insertion Sort</p>
<p>for i := 1; i < n; i++: perulangan untuk memproses setiap elemen mulai dari indeks 1</p>
<p>key := (*arr)[i]: menyimpan elemen yang akan disisipkan</p>
<p>j := i - 1: menyimpan indeks elemen sebelum key</p>
<p>for j >= 0 && (*arr)[j] > key: perulangan untuk mencari posisi yang tepat bagi key</p>
<p>(*arr)[j+1] = (*arr)[j]: menggeser elemen yang lebih besar satu posisi ke kanan</p>
<p>j--: berpindah ke indeks sebelumnya</p>
<p>(*arr)[j+1] = key: menempatkan key pada posisi yang sesuai</p>
<p>func jarak(arr tabInt, n int): fungsi untuk memeriksa apakah selisih antar data berurutan tetap</p>
<p>if n > 1: kondisi jika jumlah data lebih dari satu</p>
<p>diff := arr[1] - arr[0]: menyimpan selisih dua elemen pertama</p>
<p>status := 1: penanda bahwa selisih data masih dianggap tetap</p>
<p>for i := 1; i < n-1; i++: perulangan untuk memeriksa selisih setiap pasangan data berurutan</p>
<p>if arr[i+1]-arr[i] != diff: kondisi jika ditemukan selisih yang berbeda</p>
<p>status = 0: menandakan bahwa selisih data tidak tetap</p>
<p>if status == 1: kondisi jika semua selisih sama</p>
<p>fmt.Printf("Data berjarak %d\n", diff): menampilkan besar jarak antar data</p>
<p>else: kondisi jika ada selisih yang berbeda</p>
<p>fmt.Println("Data berjarak tidak tetap"): menampilkan bahwa jarak antar data tidak tetap</p>
<p>func main(): fungsi utama program</p>
<p>var arr tabInt: array untuk menyimpan data bilangan</p>
<p>var n int: variabel untuk menghitung jumlah data</p>
<p>var num int: variabel untuk menyimpan input</p>
<p>fmt.Scan(&num): membaca bilangan pertama</p>
<p>for num >= 0: perulangan selama input bernilai tidak negatif</p>
<p>arr[n] = num: menyimpan bilangan ke dalam array</p>
<p>n++: menambah jumlah data yang tersimpan</p>
<p>fmt.Scan(&num): membaca bilangan berikutnya</p>
<p>insertionSort(&arr, n): mengurutkan seluruh data secara menaik</p>
<p>for i := 0; i < n; i++: perulangan untuk menampilkan seluruh data yang sudah diurutkan</p>
<p>fmt.Print(arr[i]): menampilkan elemen array</p>
<p>if i < n-1: kondisi jika bukan elemen terakhir</p>
<p>fmt.Print(" "): menampilkan spasi antar data</p>
<p>fmt.Println(): pindah ke baris baru setelah semua data ditampilkan</p>
<p>jarak(arr, n): memeriksa apakah selisih antar data tetap atau tidak</p>


<p>Program ini digunakan untuk mengurutkan sekumpulan bilangan bulat dan memeriksa apakah data tersebut memiliki selisih (jarak) yang tetap antar elemen setelah diurutkan.</p>
<p>Program membaca bilangan secara berulang dan menyimpannya ke dalam array hingga pengguna memasukkan bilangan negatif sebagai tanda akhir input. Setelah seluruh data diterima, program mengurutkan data menggunakan algoritma insertion sort secara menaik.</p>
<p>Setelah proses pengurutan selesai, program menampilkan seluruh elemen array yang telah terurut dalam satu baris. Data yang sudah terurut kemudian diperiksa untuk mengetahui apakah setiap pasangan elemen yang berurutan memiliki selisih yang sama.</p>
<p>Pemeriksaan dilakukan dengan menghitung selisih antara dua elemen pertama sebagai acuan. Selanjutnya, program membandingkan selisih tersebut dengan selisih setiap pasangan elemen berikutnya. Jika seluruh selisih bernilai sama, maka data dianggap memiliki jarak tetap.</p>
<p>Pada akhir program, jika jarak antar elemen selalu sama, program menampilkan pesan "Data berjarak" beserta nilai selisihnya. Namun, jika terdapat perbedaan selisih pada salah satu pasangan elemen, program menampilkan pesan "Data berjarak tidak tetap".</p>
<p>Program tidak hanya mengurutkan data, tetapi juga menentukan apakah data tersebut membentuk suatu barisan dengan beda yang konstan.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%2014/output/soal%204.png" >

[penjelasan]
<h2>Soal 5</h2>
<br>  
<p>package main: menandakan bahwa program ini adalah program utama</p>
<p>import "fmt": digunakan untuk input/output</p>
<p>const nMax int = 7919: konstanta untuk ukuran maksimum data buku</p>
<p>type Buku struct {...}: tipe data untuk menyimpan informasi sebuah buku</p>
<p>id, judul, penulis, penerbit string: atribut identitas dan informasi buku bertipe string</p>
<p>eksemplar, tahun, rating int: atribut jumlah eksemplar, tahun terbit, dan rating bertipe integer</p>
<p>type DaftarBuku [nMax + 1]Buku: tipe data array untuk menyimpan banyak buku</p>
<p>func DaftarkanBuku(pustaka *DaftarBuku, n int): fungsi untuk mengisi data buku ke dalam array</p>
<p>for i := 1; i <= n; i++: perulangan untuk mengisi data setiap buku</p>
<p>fmt.Scan(...): membaca id, judul, penulis, penerbit, eksemplar, tahun, dan rating buku</p>
<p>func cetakFavorit(pustaka DaftarBuku, n int): fungsi untuk menampilkan buku dengan rating tertinggi</p>
<p>if n > 0: kondisi jika terdapat data buku</p>
<p>maxRating := pustaka[1].rating: menyimpan rating tertinggi sementara</p>
<p>for i := 2; i <= n; i++: perulangan untuk mencari rating tertinggi</p>
<p>if pustaka[i].rating > maxRating: kondisi jika ditemukan rating yang lebih tinggi</p>
<p>maxRating = pustaka[i].rating: memperbarui nilai rating tertinggi</p>
<p>for i := 1; i <= n; i++: perulangan untuk mencari buku dengan rating tertinggi</p>
<p>if pustaka[i].rating == maxRating: kondisi jika rating buku sama dengan rating tertinggi</p>
<p>fmt.Printf("%s, %s, %s, %d\n", ...): menampilkan judul, penulis, penerbit, dan tahun buku favorit</p>
<p>func UrutBuku(pustaka *DaftarBuku, n int): fungsi untuk mengurutkan buku berdasarkan rating secara menurun</p>
<p>for i := 2; i <= n; i++: perulangan utama pada Insertion Sort</p>
<p>key := pustaka[i]: menyimpan data buku yang akan disisipkan</p>
<p>j := i - 1: menyimpan indeks sebelum key</p>
<p>for j >= 1 && pustaka[j].rating < key.rating: perulangan untuk mencari posisi yang sesuai berdasarkan rating</p>
<p>pustaka[j+1] = pustaka[j]: menggeser data buku ke kanan</p>
<p>j--: berpindah ke indeks sebelumnya</p>
<p>pustaka[j+1] = key: menempatkan buku pada posisi yang sesuai</p>
<p>func Cetak5Terbaru(pustaka DaftarBuku, n int): fungsi untuk menampilkan maksimal 5 buku teratas setelah pengurutan</p>
<p>limit := 5: batas awal jumlah buku yang akan ditampilkan</p>
<p>if n < 5: kondisi jika jumlah buku kurang dari 5</p>
<p>limit = n: menyesuaikan jumlah buku yang ditampilkan</p>
<p>for i := 1; i <= limit; i++: perulangan untuk menampilkan buku</p>
<p>fmt.Println(pustaka[i].judul): menampilkan judul buku</p>
<p>func CariBuku(pustaka DaftarBuku, n int, r int): fungsi untuk mencari buku berdasarkan rating menggunakan Binary Search</p>
<p>kiri := 1: indeks awal pencarian</p>
<p>kanan := n: indeks akhir pencarian</p>
<p>found := -1: penanda bahwa buku belum ditemukan</p>
<p>for kiri <= kanan && found == -1: perulangan pencarian selama data belum ditemukan</p>
<p>med := (kiri + kanan) / 2: menentukan indeks tengah</p>
<p>if pustaka[med].rating == r: kondisi jika rating yang dicari ditemukan</p>
<p>found = med: menyimpan posisi buku yang ditemukan</p>
<p>else if pustaka[med].rating < r: kondisi jika rating tengah lebih kecil dari target</p>
<p>kanan = med - 1: mempersempit pencarian ke bagian kiri</p>
<p>else: kondisi jika rating tengah lebih besar dari target</p>
<p>kiri = med + 1: mempersempit pencarian ke bagian kanan</p>
<p>if found != -1: kondisi jika buku ditemukan</p>
<p>fmt.Printf("%s, %s, %s, %d, %d, %d\n", ...): menampilkan informasi lengkap buku yang ditemukan</p>
<p>else: kondisi jika buku tidak ditemukan</p>
<p>fmt.Println("Tidak ada buku dengan rating seperti itu"): menampilkan pesan bahwa buku tidak ditemukan</p>
<p>func main(): fungsi utama program</p>
<p>var n, targetRating int: variabel untuk jumlah buku dan rating yang dicari</p>
<p>var pustaka DaftarBuku: array untuk menyimpan data buku</p>
<p>fmt.Scan(&n): membaca jumlah buku</p>
<p>DaftarkanBuku(&pustaka, n): mengisi data buku ke dalam array</p>
<p>fmt.Scan(&targetRating): membaca rating yang akan dicari</p>
<p>cetakFavorit(pustaka, n): menampilkan buku dengan rating tertinggi</p>
<p>UrutBuku(&pustaka, n): mengurutkan buku berdasarkan rating dari terbesar ke terkecil</p>
<p>Cetak5Terbaru(pustaka, n): menampilkan maksimal 5 buku teratas setelah pengurutan</p>
<p>CariBuku(pustaka, n, targetRating): mencari dan menampilkan buku dengan rating tertentu</p>


<p>Program ini digunakan untuk mengelola data buku dalam sebuah perpustakaan. Setiap buku memiliki informasi berupa ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating yang disimpan dalam sebuah struktur data (struct) bernama Buku.</p>
<p>Program diawali dengan membaca jumlah buku yang akan didaftarkan, kemudian seluruh data buku dimasukkan ke dalam array DaftarBuku melalui fungsi DaftarkanBuku. Setelah itu, pengguna memasukkan sebuah nilai rating yang akan digunakan sebagai target pencarian.</p>
<p>Fungsi cetakFavorit digunakan untuk mencari rating tertinggi dari seluruh buku yang terdaftar. Setelah menemukan rating maksimum, program menampilkan informasi semua buku yang memiliki rating tersebut sehingga dapat dianggap sebagai buku favorit.</p>
<p>Selanjutnya, fungsi UrutBuku mengurutkan data buku berdasarkan rating secara menurun menggunakan algoritma insertion sort. Setelah data terurut, fungsi Cetak5Terbaru menampilkan lima buku pertama pada hasil pengurutan, atau seluruh buku jika jumlahnya kurang dari lima.</p>
<p>Juga menyediakan fitur pencarian menggunakan fungsi CariBuku. Pencarian dilakukan dengan algoritma binary search pada data yang telah diurutkan berdasarkan rating. Jika ditemukan buku dengan rating yang dicari, program akan menampilkan informasi lengkap buku tersebut. Jika tidak ditemukan, program menampilkan pesan bahwa tidak ada buku dengan rating tersebut.</p>
<p>Program ini menggabungkan proses penyimpanan data buku, pencarian buku favorit, pengurutan berdasarkan rating, penampilan lima buku dengan rating tertinggi, serta pencarian buku berdasarkan rating tertentu.</p>

##### Output
<img src = "https://github.com/leonardo0138/109082530036_Leonardo-Farriz-Garcya/blob/main/modul%2014/output/soal%205.png" >
