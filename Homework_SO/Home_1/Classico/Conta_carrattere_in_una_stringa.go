package main

import "fmt"

func main() {
	stringa := "ciaoc mondo"
	conteggioTotale := 0

	for i := 0; i < len(stringa); i++ {
		if stringa[i] == 'c' {
			conteggioTotale += 1
		}
	}
	fmt.Println(conteggioTotale)
}
