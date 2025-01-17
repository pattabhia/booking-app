package utils

import (
	fmt "fmt"
	strings "strings"
)

// PrintBookings prints the booking details and the number of bookings for the given list of booking details.
// It takes a slice of strings representing the booking details as input.
// The function extracts the first names from each booking detail and prints them along with the total number of bookings.
func PrintBookings(bookingDetails []string) {
	firstNames := []string{}
	var index int = 0
	for _, booking := range bookingDetails {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
		index++
	}
	fmt.Printf("There are %d bookings for the %s!\n", index, firstNames)
}
