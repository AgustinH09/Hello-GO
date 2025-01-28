package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	helloWorld()
	useQuote()
}

func helloWorld() {
	fmt.Println("Hello Go!")
}

func useQuote() {
	fmt.Println(quote.Go())
}
