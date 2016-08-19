package main

import (
"fmt"
)

func main(){
	i := 42
	p := &i // "=&" cria o ponteiro (atribui a variável p o endereço de memória da variável i)
	// imprime 42
	fmt.Println(*p) // "*" pega o valor de "i" através do ponteiro "p"
	// i = 42
	a := *p
	// a = 42
	i = 20
	// i = 20
	// imprime 20 e 42
	fmt.Println(*p, a)
}