package main

import "fmt"

func main() {
	// Syntatic sugar (Can't be used for constants)
	// conferenceName := "Go Conference"
	const conferenceName string = "Go Conference"
	// Go is a statically-typed language. Type can be specified, but also inferred.
	const conferenceTickets = 50
	var remainingTickets uint = 50 // Unsigned int
	// var bookings = [50]string{}
	// Alternative:
	var bookings []string // With size: array; without size: slice

	// %T for printing the type of a variable
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets -= uint(userTickets)
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole slice: %v \n", bookings)
	fmt.Printf("The first value: %v \n", bookings[0])
	fmt.Printf("Slice type: %T \n", bookings)
	fmt.Printf("Slice length: %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v. \n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v \n", bookings)
}
