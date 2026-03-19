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