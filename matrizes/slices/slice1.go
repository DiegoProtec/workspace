package main

import "fmt"

func main() {
	// slice "p" recebe uma matriz do tipo inteiro com 5(cinco) posições
	p := []int{2, 3, 5, 8, 13}
	fmt.Println("p ==", p)
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}