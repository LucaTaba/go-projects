package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var somma int = 0

	numeri := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(numeri); i++ {
		somma = somma + numeri[i]
	}

	fmt.Printf("La somma dell'array è %d \n", somma)

	elapsed := time.Since(start)
	fmt.Printf("Tempo di esecuzione: %s\n", elapsed)

}
