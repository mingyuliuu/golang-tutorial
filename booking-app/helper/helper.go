package helper

import "strings"

// Returning multiple values
func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= int(remainingTickets)

	return isValidName, isValidEmail, isValidTicketNumber
}