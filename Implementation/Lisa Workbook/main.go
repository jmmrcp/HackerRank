package main

import (
	"fmt"
)

func main() {
	var (
		n uint   // chapters
		k uint   // problems per page
		v uint   // value
		t []uint // Chapter problems
		i uint
	)

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &k) 
	t = make([]uint, n)

	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &v)
		t[i] = v
	}

	fmt.Println(t)
}
