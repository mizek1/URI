package main

import "fmt"

func main() {
	a, b, c := 0, 0, 0
	fmt.Scanf("%d %d %d\n", &a, &b, &c)
	if a >= b && a >= c {
		fmt.Println(a, "eh o maior")
	} else if b >= a && b >= c {
		fmt.Println(b, "eh o maior")
	} else if c >= a && c >= b {
		fmt.Println(c, "eh o maior")
	}
}
