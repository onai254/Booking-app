package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Nairobi DevOps Days - KCD"
var location = "Moringa"
var bookings = make([]userData, 0)

const conferenceTickers = 200

var remainingTicket uint = 200

type userData struct {
	email           string
	lastName        string
	userName        string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	//Function for greeting users

	greetUsers()

	// Main Login function

	userName, lastName, email, userTicket := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.UserInputValidation(userName, lastName, email, userTicket, remainingTicket)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTicket, userName, lastName, email)

		wg.Add(1)
		go sendTicket(userTicket, userName, lastName, email)

		// Call for function print first name

		firstNames := getFirstName()
		fmt.Printf("The First Name of the bookings are %v", firstNames)

		var ifNoTicketRemaining bool = remainingTicket == 0
		if ifNoTicketRemaining {

			// end program
			fmt.Printf("Our Conference ticket is sold out Come back nexy year \n")
			//break
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
	wg.Wait()
}
func greetUsers() {

	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("we Have a total of %v conference tickets and %v remainig \n", conferenceTickers, remainingTicket)
	fmt.Printf("Location for the Evnet is %v \n", location)
}

func getFirstName() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.userName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var userName string
	var lastName string
	var email string
	var userTicket uint

	// User Name is the name the user will create the account with

	fmt.Println("Enter Your First Name")
	fmt.Scan(&userName)

	fmt.Println("Enter Your Last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email")
	fmt.Scan(&email)

	fmt.Println("Emter Number of Tickets")
	fmt.Scan(&userTicket)

	return userName, lastName, email, userTicket
}

func bookTicket(userTicket uint, userName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTicket

	//create a map for the user
	var userData = userData{
		userName:        userName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: remainingTicket,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings %v \n", bookings)

	fmt.Printf("Thank you for booking %v Tickets You will receive your confirmation email at %v \n ", userTicket, email)
	fmt.Printf("user %v has bought %v ticket \n The remaining number of tickets are %v \n", userName, userTicket, remainingTicket)

}

func sendTicket(userTicket uint, userName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for%v %v", userTicket, userName, lastName)
	fmt.Printf("################ \n")
	fmt.Printf("Sending Ticket : \n %v \n To email address %v \n", ticket, email)
	fmt.Printf("################ \n")

	wg.Done()

}
