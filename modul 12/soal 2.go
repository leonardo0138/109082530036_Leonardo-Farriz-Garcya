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