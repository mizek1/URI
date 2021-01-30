package main

import (
	"fmt"
)

func main() {
	var n, h int
	var valor float64
	fmt.Scanf("%d\n%d\n%f", &n, &h, &valor)
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", n, (float64(h) * valor))
}
