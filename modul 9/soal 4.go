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