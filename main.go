package main

import (
	"fmt"
)

func main() {
	var conferenceName = " Nairobi DevOps Days - KCD"
	var kcdNairobi = "Kubernetes Community Days"
	var tickt = "Ticket"

	const conferenceTickers = 200

	var remainingTicket = 200

	fmt.Printf("Welcome to %v 2023 \n", conferenceName)

	fmt.Printf("we Have a total of %v conference tickets and %v remainig \n", conferenceTickers, remainingTicket)

	fmt.Printf("Its the %v Event Get your %v Here \n", kcdNairobi, tickt)

}
