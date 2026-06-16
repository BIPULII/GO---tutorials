package main

import "fmt"

func main() {
	//fmt.Println("Hello, World!");
	
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Println("Get your tickets here to attend the conference.");
	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)

	fmt.Printf("Enter your first name: ")
	var firstName string
	fmt.Scan(&firstName)
	fmt.Printf("Enter the number of tickets you want to book: ")
	var numTickets int
	fmt.Scan(&numTickets)
	fmt.Printf("Thank you %v for booking %v tickets.\n", firstName, numTickets)
}
