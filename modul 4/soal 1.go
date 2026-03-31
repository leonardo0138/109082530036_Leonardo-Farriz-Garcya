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