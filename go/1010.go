package main

import (
	"fmt"
)

func main() {
	c1, p1, v1, c2, p2, v2 := 0, 0, 0.0, 0, 0, 0.0
	fmt.Scanf("%d %d %f\n%d %d %f", &c1, &p1, &v1, &c2, &p2, &v2)
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", float64(p1)*v1+float64(p2)*v2)
}
