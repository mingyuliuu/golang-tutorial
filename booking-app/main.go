package main

import (
	"fmt"
	"strings"
)

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

	for { // 1. Infinite loop
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

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		var isValidTicketNumber bool = userTickets > 0 && userTickets <= int(remainingTickets)

		if !isValidName || !isValidEmail || !isValidTicketNumber {
			if !isValidName {
				fmt.Println("Firstname or lastname you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain the @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
			continue
		}

		remainingTickets -= uint(userTickets)
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		/*
			fmt.Printf("The first value: %v \n", bookings[0])
			fmt.Printf("Slice type: %T \n", bookings)
			fmt.Printf("Slice length: %v \n", len(bookings))
		*/

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v. \n", remainingTickets, conferenceName)

		// 2. For Each loop
		firstNames := []string{}
		for _, booking := range bookings { // Blank identifier
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all our bookings: %v \n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}
}
