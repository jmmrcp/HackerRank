package main

import "fmt"

func Multiples(N uint) {
	var suma uint = 0
    var i uint
	for i = 0; i < N; i++ {
		if i%3 == 0 || i%5 == 0 {
			suma += i
		} // End if
	} // End for
	fmt.Println(suma)
    
}

func main() {
    var (
        T uint
        N uint
        i uint
    )
    fmt.Scanf("%d", &T)
    for i = 0; i < T; i++ {
        fmt.Scanf("%d", &N)
        Multiples(N)
    }
	
}