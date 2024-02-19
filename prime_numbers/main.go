package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	ns := make([]int, n)
	for i := range ns {
		var x int
		fmt.Scanf("%d", &x)
		ns[i] = x
	}
	fmt.Println(primer(ns))
}

func primer(x []int) int {
	var count int

	for _, v := range x {
		if isPrime(v) {
			count++
		}

	}
	return count
}

func isPrime(x int) bool {
	// 2は素数なのでそのままtrue
	if x == 2 {
		return true
	}

	// 1は素数ではない
	// 2で割り切れる = 1とその数以外で割り切れるので素数ではない
	if x <= 2 || x%2 == 0 {
		return false
	}

	// 例えばx=13の場合
	// 3*3 < 13 なので、 13/3 != 0
	// 5*5 > 13 なので、forを抜けて素数
	// x=15
	// 3*3 < 15 なので 15/3 == 0 なので false
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}
