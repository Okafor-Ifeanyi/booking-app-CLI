package main

import (
	"fmt"
	"strings"
)

func main () {
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

	// Store initial variables
	conferenceName := "Ecosystem Mixer"
	const confrenceTicket int = 50 
	var remainingTicket uint = 50

	// Print welcome message
	fmt.Printf("Welcome to %v this booking service \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are stil available\n", confrenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend!")

	// Store user variables
	var bookings []string 
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for remainingTicket > 0 {
		// Getting user documents now to fill in  variables
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
 
		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter your userTickets:")
		fmt.Scan(&userTickets) 

		// input validation for user variables
		isNameValid := len(firstName) > 2 && len(lastName) > 2
		isEmailValid := strings.Contains(email, "@") || strings.Contains(email, ".")
		isTicketsValid := userTickets > 1 && userTickets < remainingTicket

		if isNameValid && isEmailValid && isTicketsValid {
			remainingTicket -= userTickets
			bookings = append(bookings, firstName + " " + lastName)
			fmt.Println(&remainingTicket, remainingTicket)

			firstNames := []string{}
			for _, booking := range bookings {
				var first = strings.Fields(booking)
				var firstName = first[1]
				firstNames = append(firstNames, firstName)
			}

			fmt.Printf("Thank you %v %v for booking %v tickets, you will recieve and email at %v shortly.\n", firstName, lastName, userTickets, email)
			fmt.Println(bookings)
			fmt.Println(firstNames)

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
				fmt.Println("Please enter a valid name, gotta be up to 2 characters")
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