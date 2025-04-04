package main

import "fmt"

func welcomeStatement() {
	fmt.Printf("Welcome to %v this booking service \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are stil available\n", confrenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend!")
}