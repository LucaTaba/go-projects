package main

func main() {
<<<<<<< HEAD
	package main

import (
	"fmt"
	"sync"
)

func calcolaSommaQuadrati(numeri []int, risultato chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Segnala che la goroutine ha finito

	somma := 0
	for _, num := range numeri {
		somma += num * num
	}
	risultato <- somma // Invia la somma parziale al canale
}

func main() {
	numeri := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numGoroutine := 3 // Numero di goroutine da utilizzare
	dimSegmento := len(numeri) / numGoroutine
	resto := len(numeri) % numGoroutine

	risultati := make(chan int, numGoroutine) // Canale per ricevere i risultati parziali
	var wg sync.WaitGroup                   // WaitGroup per attendere il completamento delle goroutine

	inizio := 0
	for i := 0; i < numGoroutine; i++ {
		fine := inizio + dimSegmento
		if i < resto {
			fine++ // Distribuisce gli elementi rimanenti tra le prime goroutine
		}
		wg.Add(1) // Incrementa il contatore del WaitGroup
		go calcolaSommaQuadrati(numeri[inizio:fine], risultati, &wg)
		inizio = fine
	}

	// Attendi che tutte le goroutine abbiano finito
	wg.Wait()
	close(risultati) // Chiudi il canale dopo che tutte le goroutine hanno inviato i loro risultati

	sommaTotale := 0
	// Ricevi i risultati parziali dal canale e calcola la somma totale
	for sommaParziale := range risultati {
		sommaTotale += sommaParziale
	}

	fmt.Println("La somma totale dei quadrati è:", sommaTotale)
}
=======

>>>>>>> 2918bfd00511d83850541a6534ba81aac75b99e1
}
