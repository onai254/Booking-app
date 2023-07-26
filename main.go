package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Nairobi DevOps Days - KCD"
	// kcdNairobi := "Kubernetes Community Days Nairobi"
	location := "Moringa"
	var bookings []string
	const conferenceTickers = 200
	var remainingTicket uint = 200

	greetUsers(conferenceName, conferenceTickers, remainingTicket, location)

	for {

		var userName string
		var userTicket uint
		var lastName string
		var email string

		// User Name is the name the user will create the account with

		fmt.Println("Enter Your First Name")
		fmt.Scan(&userName)

		fmt.Println("Enter Your Last Name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email")
		fmt.Scan(&email)

		fmt.Printf("Emter Number of Tickets")
		fmt.Scan(&userTicket)

		isValidName := len(userName) >= 6 && len(lastName) >= 6
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTicket = remainingTicket - userTicket
			bookings = append(bookings, userName+" "+lastName)

			fmt.Printf("Thank you for booking %v Tickets You will receive your confirmation email at %v \n ", userTicket, email)
			fmt.Printf("user %v has bought %v ticket \n The remaining number of tickets are %v \n", userName, userTicket, remainingTicket)

			// First Name The first Name of the users will be displayed when the user finishes registering.

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first name of bookings: %v \n", firstNames)

			var ifNoTicketRemaining bool = remainingTicket == 0
			if ifNoTicketRemaining {

				// end program
				fmt.Printf("Our Conference ticket is sold out Come back nexy year \n")
				break
			}

		} else {
			if !isValidEmail {
				fmt.Print("Check the Email Does not have the @ match \n")
			}
			if !isValidName {
				fmt.Print("Check the First NAme and Last Name does not contain 6 characters \n")
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets are invalid\n")
			}
		}
	}
}
func greetUsers(conferenceName string, conferenceTickers int, remainingTicket uint, location string) {

	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("we Have a total of %v conference tickets and %v remainig \n", conferenceTickers, remainingTicket)
	fmt.Printf("Location for the Evnet is %v \n", location)
}
