package main

import "fmt"

func main() {
	var a, i, j uint32
	var b, dia1, dia2, res int32

	fmt.Scan(&a)

	arr := make([][]int32, a)
	for i = 0; i < a; i++ {
		arr[i] = make([]int32, a)
	}
	for i = 0; i < a; i++ {
		for j = 0; j < a; j++ {
			fmt.Scan(&b)
			arr[i][j] = b
		}
	}
	j = a - 1
	for i = 0; i < a; i++ {
		dia1 += arr[i][i]
		dia2 += arr[i][j-i]
	}
	res = dia1 - dia2
	if res < 0 {
		res = -res
	}
	fmt.Printf("%v\n", res)
}
