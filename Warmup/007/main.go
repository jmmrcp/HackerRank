package main

import (
	"fmt"
	"strconv"
)

func militar(hora string) string {
	// hh:mm:ssAM or hh:mm:ssPM //
	var (
		h string // horas
		t string // tipo AM o PM
		n int    // hora en entero
	)
	h = hora[:2]
	t = hora[8:10]
	if h == "12" && t == "AM" {
		h = "00"
	}
	if h == "12" && t == "PM" {
		h = "12"
	}
	if h != "12" && t == "PM" {
		n, _ = strconv.Atoi(h)
		n += 12
		h = strconv.Itoa(n)
	}
	return h + ":" + hora[3:8]
}
func main() {
	var (
		a string // Hora a convertir a militar
	)
	fmt.Scan(&a)
	fmt.Println(militar(a))
}
