package main

import "fmt"

func main() {
	a := make([]int, 0, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	//capacidade = capacidade da matrizb =5, - 0, "b[:"
	c := b[:2]
	printSlice("c", c)
	//capacidade = capacidade da matrizb =5, - 2 "c[2:"
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int){
	fmt.Printf("%s len == %d, cap == %d, %v\n", s, len(x), cap(x), x)
}