package main

import (
	"fmt"
	"regexp"
	"strings"
)

var conferenceName string = "GopherCon"

const conferenceTickets int = 50

var bookingDetails = []string{}
var userTickets int
var remainingTickets = conferenceTickets
var firstName string
var lastName string

func main() {

	for {
		if remainingTickets == 0 {
			fmt.Println("Sorry, all tickets have been sold out!")
			break
		}

		getUserInputs(firstName, lastName, userTickets)
		printRemainingTickets()
		printBookings(bookingDetails)
	}

}

func getUserInputs(firstName string, lastName string, userTickets int) {

	fmt.Println("Please enter your firstName:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your lastName:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter the number of tickets you want to buy:")
	fmt.Scan(&userTickets)

	isValidInputs := validateUserInputs(firstName, lastName, userTickets)

	if isValidInputs {
		bookTickets(firstName, lastName, userTickets)
	}
}

func bookTickets(firstName, lastName string, userTickets int) {
	bookingDetails = append(bookingDetails, firstName+" "+lastName)
	remainingTickets = calculateRemainingTickets(userTickets)
	fmt.Printf("Thank you, %s! You have bought %d tickets for the %s!\n", firstName, userTickets, conferenceName)
}

func validateUserInputs(firstName, lastName string, userTickets int) bool {
	switch {
	case len(firstName) == 0 || len(lastName) == 0:
		fmt.Println("Names cannot be empty!")
		return false
	case !regexp.MustCompile(`^[a-zA-z]+$`).MatchString(firstName) ||
		!regexp.MustCompile(`^[a-zA-z]+$`).MatchString(lastName):
		fmt.Println("Names cannot contain numbers!")
		return false
	case userTickets <= 0:
		fmt.Println("Number of tickets should be greater than 0!")
		return false
	case userTickets > remainingTickets:
		fmt.Println("Sorry, there are not enough tickets available!")
		return false
	default:
		return true
	}
}

func calculateRemainingTickets(userTickets int) int {
	return remainingTickets - userTickets
}

func printRemainingTickets() {
	fmt.Printf("There are %d tickets remaining for the %s!\n", remainingTickets, conferenceName)
}

func printBookings(bookingDetails []string) {
	firstNames := []string{}
	var index int = 0
	for _, booking := range bookingDetails {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
		index++
	}
	fmt.Printf("There are %d bookings for the %s!\n", index, firstNames)

	titles := []string{"Mr", "Mrs", "Ms", "Dr", "Prof"}
	for _, title := range titles {
		fmt.Printf("%s \n", title)
	}
}
