package main

import "fmt"

func main() {
	const Pi = 3.14159
	r := 0.0
	fmt.Scanf("%f", &r)
	fmt.Printf("VOLUME = %.3f\n", ((4 / 3.0) * Pi * (r * r * r)))
}
