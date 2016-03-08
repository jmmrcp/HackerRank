package main

import "fmt"

func kaprekar(n uint64) bool {
	order := 0
	if n == 1 {
		return true
	}
	nn, power := n*n, uint64(1)
	for power <= nn {
		power *= 10
		order++
	}
	power /= 10
	order--
	for ; power > 1; power /= 10 {
		q, r := nn/power, nn%power
		if q >= n {
			return false
		}
		if q+r == n {
			return true
		}
		order--
	}
	return false
}
func main() {
	var (
		p uint64
		q uint64
	)
	fmt.Scan(&p)
	fmt.Scan(&q)

	for p < q {
		is := kaprekar(p)
		if is {
			fmt.Printf("%v ", p)
		}
		p++
	}
}
