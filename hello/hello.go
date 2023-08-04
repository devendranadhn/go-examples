package main

import (
	"fmt"

	"rsc.io/quote"

	"github.com/devendranadhn/greetings"
)

func main() {

	fmt.Println("Hello, world")

	fmt.Println(quote.Go())

	message := greetings.Hello("Deva")
	fmt.Println(message)

}
