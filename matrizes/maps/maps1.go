package main

import "fmt"

type Vetor struct {
	Lat, Long float64
}

func main() {

	m := make(map[string]Vetor)
	m["LabDiego"] = Vetor{
		40.68433, -74.39967,
	}
	fmt.Println(m["LabDiego"])
}