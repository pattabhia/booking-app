package utils

import (
	fmt "fmt"
	strings "strings"
)

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
