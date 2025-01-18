package main

import (
	"booking-app/service"
	"fmt"
)

var conferenceName string = "GopherCon"

const conferenceTickets int = 50

var userTickets int
var firstName string
var lastName string

func main() {

	for {
		if service.GetRemainingTickets() == 0 {
			break
		}
		firstName, lastName, userTickets = getUserInputs(firstName, lastName, userTickets)
		isValidInputs := service.ValidateUserInputs(firstName, lastName, userTickets)
		if isValidInputs {
			service.BookTickets(firstName, lastName, userTickets)
		}
	}
}

func getUserInputs(firstName string, lastName string, userTickets int) (string, string, int) {

	fmt.Println("Please enter your firstName:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your lastName:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter the number of tickets you want to buy:")
	fmt.Scan(&userTickets)

	return firstName, lastName, userTickets
}
