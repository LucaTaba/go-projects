package main

import "fmt"

func main() {
	var nome string
	var eta int

	n, err := fmt.Scanln(&nome, &eta)
	fmt.Println("Ciao, mondo!")

	if err != nil {
		fmt.Println("Errore durante la lettura:", err)
	}
	fmt.Println("Errore durante la lettura:", n)
	return

}
