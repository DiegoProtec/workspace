package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 8, 13, 21}
	fmt.Println("p ==", p)
	// mantém após o índice 1 e até o índice 4
	fmt.Println("p[1:4] ==", p[1:4])
	// mantém até o índice 3
	fmt.Println("p[:3] ==", p[:3])
	// mantém após o índice 4
	fmt.Println("p[4:] ==", p[4:])
}