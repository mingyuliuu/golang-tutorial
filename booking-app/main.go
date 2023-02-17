package main

import "fmt"

func main() {
	// Syntatic sugar (Can't be used for constants)
	// conferenceName := "Go Conference"
	const conferenceName string = "Go Conference"
	// Go is a statically-typed language. Type can be specified, but also inferred.
	const conferenceTickets = 50
	var remainingTickets uint = 50 // Unsigned int

	// %T for printing the type of a variable
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
