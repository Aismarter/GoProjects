package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main() {
	// Set properties of the predefined Logger, incluing
	// the log entry prefix and a flag to disable printing
	// the time, source code, and liner number.
	log.SetPrefix("greeting: ")
	log.SetFlags(0)
	
	// Get a greeting message and print it.
	//message := greetings.Hello("Gladys")
	
	// Request a greeting message.
	message, err := greetings.Hello("")
	// If an erro was returned, print it to the console and
	//exit the program
	if err != nil{
		log.Fatal(err)
	}
	
	// If no error was returned, print the returned message
	// to the console
	fmt.Println(message)
}
