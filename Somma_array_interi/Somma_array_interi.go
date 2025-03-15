package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	//gestione input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	//inizio tempo
	start := time.Now()

	input := scanner.Text()

	numeriStringhe := strings.Split(input, " ")

	numeri := make([]int, len(numeriStringhe))

	// Itera sulle sottostringhe e convertile in interi.
	for i, numStr := range numeriStringhe {
		// Usa strconv.Atoi per convertire la sottostringa in un intero.
		num, err := strconv.Atoi(numStr)
		if err != nil {
			// Gestisci l'errore se la conversione fallisce.
			fmt.Println("Errore di conversione:", err)
			return
		}
		// Assegna l'intero all'array.
		numeri[i] = num
	}

	var somma int = 0

	for i := 0; i < len(numeri); i++ {
		somma = somma + numeri[i]
	}

	fmt.Printf("La somma dell'array è %d \n", somma)

	elapsed := time.Since(start)
	fmt.Printf("Tempo di esecuzione: %s\n", elapsed)

}
