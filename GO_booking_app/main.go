package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50 // total number of tickets

var conferenceName = "Go Conference" // conference name
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // UserData is a struct that holds user data

type UserData struct { //declare a struct
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{} //declare a wait group very important

func main() {

	greetUsers() //declare a function

	// for {
	firstName, lastName, email, userTickets := getUserInput() //declare a function
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber { //check if all the conditions are true

		bookTicket(userTickets, firstName, lastName, email) //book ticket

		wg.Add(1) //increment the counter
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.") //print a message
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	//}
	wg.Wait() //wait for all goroutines to finish
}

func greetUsers() { //greet users function
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string { //get first names function
	firstNames := []string{}           //declare a slice of strings
	for _, booking := range bookings { //loop through bookings
		firstNames = append(firstNames, booking.firstName) //append first name to firstNames slice
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

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

func bookTicket(userTickets uint, firstName string, lastName string, email string) { //book ticket function
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) { //send ticket function
	time.Sleep(10 * time.Second)                                                       //sleep for 50 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //declare a string
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done() //decrement the counter
}

/* var bookings = make([]map[string]string, 0) // UserData is a map that holds user data
// create user map
var user = make(map[string]string) //create a map of strings
user["firstName"] = firstName
user["lastName"] = lastName
user["email"] = email
user["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) */ //convert userTickets to string
