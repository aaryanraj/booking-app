package services

import (
	"fmt"

	"github.com/aaryanraj/booking-app/pkg/structs"
)

func CheckAvalability(remTickets *structs.Availability) {
	/*
		remainingTickets := "X"
		fmt.Printf("We have a total of %v tickets and %v are remaining\n", remainingTickets, remainingTickets)
		fmt.Printf("%v tickets are remaining for the %v\n", remainingTickets, constants.ConferenceName)
	*/

	fmt.Println("Ticket Avalibility in different Cities are as follows:")
	fmt.Printf("Mumbai : %v\n", remTickets.Mumbai)
	fmt.Printf("Bengaluru : %v\n", remTickets.Bengaluru)
	fmt.Printf("Delhi : %v\n", remTickets.Delhi)
	fmt.Printf("Kolkata : %v\n", remTickets.Kolkata)
	fmt.Println("Enter any key to return to main menu.")
	var input string
	fmt.Scan(&input) //needs work

}
