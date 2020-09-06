package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func test() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Leo",
		"Gui",
		"Omar",
	}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
