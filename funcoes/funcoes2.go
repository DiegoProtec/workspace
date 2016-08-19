package main

import "fmt"

func adder() func(int) int {
	num := 0
	return func(x int) int{
		num += x
		return num
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}