package service

import (
	"booking-app/utils"
	"fmt"
	"regexp"
	"strconv"
)

const conferenceTickets int = 50

var remainingTickets = conferenceTickets
var bookings = make([]map[string]string, 0)

func ValidateUserInputs(firstName, lastName string, userTickets int) bool {
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
		fmt.Printf("There are only %d tickets remaining!\n", remainingTickets)
		return false
	default:
		return true
	}
}

func BookTickets(firstName, lastName string, userTickets int) {
	if remainingTickets == 0 {
		fmt.Println("Sorry, there are no more tickets available!")
		return
	}
	var userMap = make(map[string]string)
	userMap["firstName"] = firstName
	userMap["lastName"] = lastName
	userMap["ticketsCount"] = strconv.FormatInt(int64(userTickets), 10)

	for key, value := range userMap {
		fmt.Printf("Key: %s, Value:%s \n", key, value)
	}
	bookings = append(bookings, userMap)
	remainingTickets = calculateRemainingTickets(userTickets, remainingTickets)
	fmt.Printf("Thank you, %s! You have bought %d tickets!\n", firstName, userTickets)
	fmt.Printf("There are %d tickets remaining!\n", remainingTickets)
	utils.PrintBookings(bookings)

}

func calculateRemainingTickets(userTickets int, remainingTickets int) int {
	return remainingTickets - userTickets
}

func GetRemainingTickets() int {
	return remainingTickets
}
