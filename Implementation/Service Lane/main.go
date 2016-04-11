package main

import "fmt"

// Min ...
func Min(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}

func vehicle(x, y uint, s []uint) {
	var min uint = 999
	for i := x; i <= y; i++ {
		min = Min(min, s[i])
	}
	fmt.Println(min)
}
func main() {
	var (
		t uint // Numero de elementos del Array
		n uint // Numero de segmentos a comparar
		i uint // variable para el bucle
		v uint // valor en el array

		inicio uint
		final  uint
	)

	fmt.Scanf("%d", &t)
	fmt.Scanf("%d", &n)
	slice := make([]uint, t)

	for i = 0; i < t; i++ {
		fmt.Scanf("%d", &v)
		slice[i] = v
	}
	// fmt.Println("Valor del Slice", slice)

	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &inicio)
		fmt.Scanf("%d", &final)

		vehicle(inicio, final, slice)
	}
}
