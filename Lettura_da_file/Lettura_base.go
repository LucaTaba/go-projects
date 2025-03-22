package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" // Pacchetto per convertire stringhe in numeri
)

func main() {
	file, err := os.Open("C:\\Users\\tabar\\OneDrive\\Desktop\\GitHub\\go-projects\\Lettura_da_file\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // con questa prima parte andiamo a leggere dal file input (in sola lettura) tramite il pacchetto os

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numero, err := strconv.Atoi(line) // Converte la stringa in un intero
		if err != nil {
			log.Printf("Errore di conversione: %s\n", line)
			continue // Salta alla riga successiva in caso di errore
		}
		fmt.Println(numero)
	}

	if err := scanner.Err(); err != nil {

		log.Fatal(err)
	}
}
