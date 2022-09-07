package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Gladys")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	// Request a greeting message, with an empty name.
	message, err = greetings.Hello("")
	
	// This will generate an error! 
	// Print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console. -- this code will never execute!
	fmt.Println(message)
}
