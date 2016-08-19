package main

import "fmt"

func main() {
	// matriz, de uma linha e duas colunas, de tipo string
	var a [2]string
	a[0] = "Olá"
	a[1] = "mundo"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}