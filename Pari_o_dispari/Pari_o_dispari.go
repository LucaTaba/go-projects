package main

import "fmt"

func main() {
	var numero int

	n, err := fmt.Scanln(&numero)
	fmt.Println("Ciao, mondo!")

	if err != nil {
		fmt.Println("Errore durante la lettura:", err)
	}
	fmt.Println("Il tuo numero è:", numero)
	fmt.Println("Elementi letti:", n)
	return

}
