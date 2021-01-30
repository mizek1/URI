package main

import (
	"fmt"
)

func main() {
	var nome string
	var salFixo, totalVendas float64
	fmt.Scanf("%s\n%f\n%f", &nome, &salFixo, &totalVendas)
	fmt.Printf("TOTAL = R$ %.2f\n", (salFixo + (totalVendas * 0.15)))
}
