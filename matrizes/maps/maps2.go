package main

import "fmt"

type Vetor struct {
	Lat, Long float64
}

var m = map[string]Vetor{
	"BelLabs": Vetor{
		40.68433, -74.39967,
	},
	"DiegoLabs": Vetor{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}