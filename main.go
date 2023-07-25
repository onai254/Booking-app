package main

import (
	"fmt"
)

func main() {
	conferenceName := " Nairobi DevOps Days - KCD"
	var kcdNairobi = "Kubernetes Community Days"
	var ticket = "Ticket"
	var location = "Moringa"

	const conferenceTickers = 200

	var remainingTicket uint = 200

	fmt.Printf("Conference Name is which type %T \n The Ticket is chich type %T \n the Location is which type %T \n", conferenceName, conferenceTickers, location)

	fmt.Printf("Welcome to %v 2023 \n", conferenceName)

	fmt.Printf("we Have a total of %v conference tickets and %v remainig \n", conferenceTickers, remainingTicket)

	fmt.Printf("Its the %v Event Get your %v Here \n", kcdNairobi, ticket)

	fmt.Println("The event will be hosted at ", location)

	var userName string
	var userTicker int

	// User Name
	fmt.Scan(&userName)

	userTicker = 2

	fmt.Printf("user %v has bought %v ticket \n ", userName, userTicker)

}
