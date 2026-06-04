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