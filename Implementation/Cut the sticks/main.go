package main

import "fmt"

func Min(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}

func Cut(s []uint) uint {
	var min uint = 999
	for _, n := range s {
		min = Min(min, n)
	}
	return min
}

func Sticks(s []uint) []uint {
	var a []uint
	if len(s) > 0 {
		n := Cut(s)
		for _, v := range s {
			d := v - n
			if d > 0 {
				a = append(a, d)
			}
		}
		fmt.Println(len(s))
		return Sticks(a)
	}
	return a
}

func main() {

	var N uint // numero de sticks
	var l uint // longitud del stick
	var i uint

	fmt.Scanf("%d", &N)
	a := make([]uint, N)

	for i = 0; i < N; i++ {
		fmt.Scanf("%d", &l)
		a[i] = l
	}
	Sticks(a)
}
