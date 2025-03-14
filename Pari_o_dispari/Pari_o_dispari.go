package main

import "fmt"

func main() {
	var numero int

	n, err := fmt.Scanln(&numero)

	if err != nil {
		fmt.Println("Errore durante la lettura:", err)
	}
	if numero%2 == 0 {
		fmt.Println("Il tuo numero è pari")
	} else {
		fmt.Println("Il tuo numero è dispari")
	}
	fmt.Println("Elementi letti:", n)
	return

}
