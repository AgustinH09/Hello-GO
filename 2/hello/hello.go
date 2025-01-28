package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Agustin")
	if err != nil {
		log.Fatal(err)
	}
	messages, err := greetings.Hellos(
		[]string{"Agustin", "Camila", "Alexis", "ThePrimagen"},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	fmt.Println(message)
}
