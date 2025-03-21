package main

import (
	"fmt"
	"sync"
)

func sommaParziale(numeri []int, risultato chan int, wg *sync.WaitGroup) {
	defer wg.Done() //Indica che la goroutine ha terminato
	somma := 0
	for _, num := range numeri {
		somma += num
	}
	risultato <- somma
}

func main() {
	numeri := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	dimensioneParte := 5
	numGoroutine := len(numeri) / dimensioneParte

	risultato := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < numGoroutine; i++ {
		inizio := i * dimensioneParte
		fine := inizio + dimensioneParte
		wg.Add(1) //incrementa il contatore delle goroutine
		go sommaParziale(numeri[inizio:fine], risultato, &wg)
	}

	go func() {
		wg.Wait() //Attende che tutte le goroutine siano terminate
		close(risultato)
	}()

	sommaTotale := 0
	for parziale := range risultato {
		sommaTotale += parziale
	}

	fmt.Println("Somma totale:", sommaTotale)

}
