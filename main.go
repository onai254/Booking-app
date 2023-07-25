package main

import (
	"fmt"
)

func main() {
	conferenceName := " Nairobi DevOps Days - KCD"
	var kcdNairobi = "Kubernetes Community Days"
	var ticket = "Ticket"
	var location = "Moringa"
	var bookings []string

	const conferenceTickers = 200

	var remainingTicket uint = 200

	fmt.Printf("Welcome to %v 2023 \n", conferenceName)

	fmt.Printf("we Have a total of %v conference tickets and %v remainig \n", conferenceTickers, remainingTicket)

	fmt.Printf("Its the %v Event Get your %v Here \n", kcdNairobi, ticket)

	fmt.Println("The event will be hosted at ", location)

	for {

		var userName string
		var userTicket uint
		var lastName string
		var email string

		// User Name

		fmt.Println("Enter Your First Name")
		fmt.Scan(&userName)

		fmt.Println("Enter Your Last Name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email")
		fmt.Scan(&email)

		fmt.Printf("Emter Number of Tickets")
		fmt.Scan(&userTicket)

		remainingTicket = remainingTicket - userTicket

		bookings = append(bookings, userName+" "+lastName)

		fmt.Printf("Thank you for booking %v Tickets You will receive your confirmation email at %v \n ", userTicket, email)

		fmt.Printf("user %v has bought %v ticket \n The remaining number of tickets are %v", userName, userTicket, remainingTicket)

		fmt.Printf("This are all our bookings: %v", bookings)

	}

}
