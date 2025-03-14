package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Ciao, mondo!")

	elapsed := time.Since(start)
	fmt.Printf("Tempo di esecuzione: %s\n", elapsed)
}
