package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukan 4 bilangan asli (a b c d): ")
	fmt.Scanln(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println("Keluaran:")
		fmt.Printf("Permutasi dan Kombinasi a terhadap c: %d %d\n", permutasi(a, c), kombinasi(a, c))
		fmt.Printf("Permutasi dan Kombinasi b terhadap d: %d %d\n", permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Println("Syarat tidak terpenuhi. a harus lebih besar atau sama dengan c, dan b harus lebih besar atau sama dengan d.")
	}
}
