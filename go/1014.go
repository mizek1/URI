package main

import "fmt"

func main() {
	a, b := 0, 0.0
	fmt.Scanf("%d\n%f", &a, &b)
	fmt.Printf("%.3f km/l\n", float64(a)/b)
}
