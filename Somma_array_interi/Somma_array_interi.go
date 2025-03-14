package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	numeri := [5]int{1, 2, 3, 4, 5}

	fmt.Println(numeri[0])

	elapsed := time.Since(start)
	fmt.Printf("Tempo di esecuzione: %s\n", elapsed)
}
