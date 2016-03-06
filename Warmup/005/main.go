package main

import "fmt"

func main() {
	var (
		a       uint32 // tamaño del Array
		b       int32  // entrada de datos ...parametros
		i       uint32 // variable para for > 0
		contpos uint32 // contador de positivos
		contneg uint32 // contador de negativos
		contcer uint32 // contador ceros
	)

	fmt.Scan(&a)

	for i = 0; i < a; i++ {
		fmt.Scan(&b)
		if b == 0 {
			contcer++
		}
		if b > 0 {
			contpos++
		}
		if b < 0 {
			contneg++
		}
	}
	fmt.Printf("%0.06v\n", float64(contpos)/float64(a))
	fmt.Printf("%0.06v\n", float64(contneg)/float64(a))
	fmt.Printf("%0.06v\n", float64(contcer)/float64(a))
}
