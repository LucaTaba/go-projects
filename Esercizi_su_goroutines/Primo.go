package main

import (
	"fmt"
	"time"
)

func stampaNumeri() {
	for i := 1; i <= 8; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func stampaLettere() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Println(string(i))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go stampaNumeri()
	go stampaLettere()

	// Aspetta un po' per far completare le goroutine
	time.Sleep(12 * time.Second)
}
