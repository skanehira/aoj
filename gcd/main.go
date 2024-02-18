package main

import "fmt"

func main() {
	var a, b int64
	fmt.Scanf("%d %d", &a, &b)

	if a < b {
		a, b = b, a
	}

	fmt.Println(gcd(a, b))
}

func gcd(a int64, b int64) int64 {
	var r int64
	for {
		r = a % b
		if r == 0 {
			break
		}

		a = b
		b = r
	}
	return b
}
