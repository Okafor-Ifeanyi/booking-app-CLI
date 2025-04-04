package helper

import "fmt"

func InputData() (string, string, string, uint) {
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