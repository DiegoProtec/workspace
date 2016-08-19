package main

import "fmt"

type Vetor struct {
	X, Y int
	B bool
}

var (
	// uma struct literal, denota valor default para todos os campos que não foram invocados
	v1 = Vetor{}
	//v2.X = 3, v2.B = false e v2.X = 3
	v2 = Vetor{X: 3}
	// caso não seja struct literal, é necessário inicializar todos os campos
	v3 = Vetor{1, 2, true}
	// o prefixo & constrói um ponteiro para uma struct literal
	p = &Vetor{Y: 4}
)
func main() {
	// para acessar os campos da struct basta usar o prefixo "."
	fmt.Println(v1, v2, v2.X, v2.Y, v2.B, v3, v3.B, *p)
}