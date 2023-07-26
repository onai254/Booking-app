package helper

import "strings"

func UserInputValidation(userName string, lastName string, email string, userTicket uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(userName) >= 6 && len(lastName) >= 6
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

	return isValidName, isValidEmail, isValidTicketNumber
}
