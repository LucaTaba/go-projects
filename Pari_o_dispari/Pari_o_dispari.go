package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Inserisci un numero:")
	n, err := fmt.Scanln(&numero)

	if err != nil {
		fmt.Println("Errore durante la lettura:", err)
	}
	if numero%2 == 0 {
		fmt.Printf("Il numero %d è pari.\n", numero)
	} else {
		fmt.Printf("Il numero %d è dispari.\n", numero)
	}
	fmt.Printf("Elementi letti: %d\n", n)
	return

}
