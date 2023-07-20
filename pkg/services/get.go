package services

import (
	"fmt"

	"github.com/aaryanraj/booking-app/pkg/structs"
)

func Get(bookings map[string][]structs.UserDataModel) {
	var id, city string
	var details structs.UserDataModel
	var bookingfound bool
	fmt.Println("Enter the Ticket ID: ")
	fmt.Scan(&id)

	fmt.Println("Enter City: ")
	fmt.Scan(&city)

	if value, found := bookings[city]; found {
		for _, booking := range value {
			if booking.BookingID == id {
				details = booking
				bookingfound = true
				break
			}
		}
	}

	if bookingfound {
		fmt.Printf("\nYou have a booking with US\n")
		fmt.Printf("Booking ID: %v\n", details.BookingID)
		fmt.Printf("Name: %v %v\n", details.FirstName, details.LastName)
		fmt.Printf("Number of Tickets: %v\n", details.UserTickets)
		fmt.Printf("Email: %v\n", details.Email)
		fmt.Printf("City: %v\n", details.City)
	} else {
		fmt.Println("\nTicket Not Found")
		fmt.Println("Details did not match. please try again with correct details!")
	}

	fmt.Println("Enter any key to return to main menu.")
	var input string
	fmt.Scan(&input) //needs work
}
