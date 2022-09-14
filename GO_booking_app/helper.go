package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) { //
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")                              // check if email contains @
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // check if userTickets is greater than 0 and less than remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
