package utils

import (
	fmt "fmt"
	"strings"
)

func PrintBookings(bookings []map[string]string) {
	firstNames := []string{}
	lastNames := []string{}
	ticketsCounts := []string{}
	var index int = 0
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
		lastNames = append(lastNames, booking["lastName"])
		ticketsCounts = append(ticketsCounts, booking["ticketsCount"])
		index++
	}
	fmt.Printf("There are %d bookings for the %s %s having %s!\n", index, firstNames, lastNames, strings.Join(ticketsCounts, ", "))
}
