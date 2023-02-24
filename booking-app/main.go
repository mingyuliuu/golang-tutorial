package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Syntatic sugar (Can't be used for constants)
// conferenceName := "Go Conference" (can't be used for package level variables)
const conferenceName string = "Go Conference"

// Go is a statically-typed language. Type can be specified, but also inferred.
const conferenceTickets = 50

var remainingTickets uint = 50 // Unsigned int
// Alternative: var bookings = [50]string{}
// With size: array; without size: slice
var bookings = make([]UserData, 0) // Initialize a slice of maps

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for { // 1. Infinite loop
		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

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

		bookTicket(uint(userTickets), firstName, lastName, email)

		wg.Add(1) // Sets the number of goroutines to wait for
		go sendTicket(uint(userTickets), firstName, lastName, email)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}

	wg.Wait() // Blocks until the WaitGroup counter is 0
}

func greetUsers() {
	// %T for printing the type of a variable
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func printFirstNames() {
	// 2. For Each loop
	firstNames := []string{}
	for _, booking := range bookings { // Blank identifier
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("These are all our bookings: %v \n", firstNames)
}

func getUserInputs() (string, string, string, int) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is: %v \n", bookings)

	/*
		fmt.Printf("The first value: %v \n", bookings[0])
		fmt.Printf("Slice type: %T \n", bookings)
		fmt.Printf("Slice length: %v \n", len(bookings))
	*/

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v. \n", remainingTickets, conferenceName)
	printFirstNames()
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// The "sleep" function stops or blocks the current "thread" (goroutine) execution for the defined duration
	time.Sleep(5 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket: \n%v to email address %v. \n", ticket, email)
	fmt.Println("###############")

	wg.Done() // Decrements the WaitGroup counter by 1
}
