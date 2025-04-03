package main

import (
	"fmt"
	"strings"
)

// Store initial variables
const confrenceTicket int = 50 
var conferenceName = "Ecosystem Mixer"
var remainingTicket uint = 50
var bookings []string 

func main () {
	// Print welcome message
	welcomeStatement()

	for remainingTicket > 0 {
		// Getting user documents now to fill in variables
		firstName, lastName, email, userTickets := inputData()

		// input validation for user variables
		isNameValid, isEmailValid, isTicketsValid := validateInputs(firstName, lastName, email, userTickets, remainingTicket)
		
		if isNameValid && isEmailValid && isTicketsValid {
			// book tickets from the remaining
			bookTicket(remainingTicket, userTickets, bookings, firstName, lastName, email)

			if remainingTicket == 0 {
				fmt.Println("We are sold out")
				break
			}
		} else if userTickets == remainingTicket {
			// If the user tries to book more tickets than available
			fmt.Println("You got the last tickets")
			break
		} else {
			if !isNameValid {
				fmt.Println("Please enter a valid name, gotta be more than 2 characters")
			}
			if !isEmailValid {
				fmt.Println("Invalid email, please enter a valid email")
			}
			if !isTicketsValid { 
				fmt.Printf("Sorry, we only have %v tickets remaining\n Try a valid number", remainingTicket)
			}
			continue
		}
	}
}


func welcomeStatement() {
	fmt.Printf("Welcome to %v this booking service \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are stil available\n", confrenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames( bookings []string ) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var first = strings.Fields(booking)
		var firstName = first[0]
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}

func validateInputs(firstName string, lastName string, email string, userTickets uint, remainingTicket uint) (bool, bool, bool) {
	isNameValid := len(firstName) > 2 && len(lastName) > 2
	isEmailValid := strings.Contains(email, "@") && strings.Contains(email, ".")
	isTicketsValid := userTickets > 0 && userTickets < remainingTicket

	return isNameValid, isEmailValid, isTicketsValid
}

func inputData() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Getting user documents now to fill in  variables
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter your userTickets:")
	fmt.Scan(&userTickets) 

	return firstName, lastName, email, userTickets 
}

func bookTicket(remainingTicket uint, userTickets uint, bookings []string, firstName string, lastName string, email string) {
	remainingTicket -= userTickets
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Println(&remainingTicket, remainingTicket)

	firstnames := getFirstNames(bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will recieve and email at %v shortly.\n", firstName, lastName, userTickets, email)
	fmt.Println(bookings)
	fmt.Printf("The first names of the bookings are: %v\n", firstnames)
}


	// fmt.Println("Welcome to our confrence booking application")
	// fmt.Println("Get your ticekts to attend!")

	// const name = "John Doe"
	// var age = 23
	
	// fmt.Println("Hello", name, "of age", age)

	// var tickets int

	// tickets = age + 10
	// fmt.Printf("Tickets available: %v \n", tickets)

	// var last_name string
	// fmt.Print("Enter your last name: ")
	// fmt.Scan(&last_name)

	// fmt.Println("Hello", name, last_name)